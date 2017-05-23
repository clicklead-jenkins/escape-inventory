/*
Copyright 2017 Ankyra

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"bytes"
	"encoding/json"
	"github.com/ankyra/escape-registry/cmd"
	"github.com/ankyra/escape-registry/dao"
	. "gopkg.in/check.v1"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type suite struct{}

var _ = Suite(&suite{})

var handler http.Handler
var rr *httptest.ResponseRecorder

func (s *suite) SetUpTest(c *C) {
	dao.TestSetup()
	handler = cmd.GetHandler(getMux())
	rr = httptest.NewRecorder()
}

const (
	registerEndpoint = "/a/my-project/register"

	applicationsTestProject       = "applications-test"
	applicationsEndpoint          = "/a/" + applicationsTestProject + "/"
	applicationsEndpointNoProject = "/a/doesnt-exist/"

	applicationVersionsTestProject = "versions-test"
	applicationVersionsEndpoint    = "/a/" + applicationVersionsTestProject + "/my-app/"
	applicationVersionsNoProject   = "/a/doesnt-exist/my-app/"
	applicationVersionsNoApp       = "/a/versions-test/doesnt-exist/"
)

func testRequest(c *C, req *http.Request, expectedStatus int) {
	handler.ServeHTTP(rr, req)
	c.Assert(rr.Code, DeepEquals, expectedStatus, Commentf("Responded with body '%s'", rr.Body.String()))
}

func (s *suite) addRelease(c *C, project, version string) {
	body := bytes.NewReader([]byte(`{"name": "my-app", "version": "` + version + `"}`))
	req, _ := http.NewRequest("POST", "/a/"+project+"/register", body)
	testRequest(c, req, 200)
	rr = httptest.NewRecorder()
}

func (s *suite) Test_Register_fails_with_empty_body(c *C) {
	req, _ := http.NewRequest("POST", registerEndpoint, nil)
	testRequest(c, req, 400)
}

func (s *suite) Test_Register_fails_with_invalid_json(c *C) {
	body := bytes.NewReader([]byte("hello"))
	req, _ := http.NewRequest("POST", registerEndpoint, body)
	testRequest(c, req, 400)
}

func (s *suite) Test_Register_fails_if_required_fields_are_missing(c *C) {
	cases := []string{
		`{"name": "missing-version"}`,
		`{"version": "1"}`,
		`{}`,
		`[]`,
		`null`,
		`12`,
	}
	for _, testCase := range cases {
		body := bytes.NewReader([]byte(testCase))
		req, _ := http.NewRequest("POST", registerEndpoint, body)
		testRequest(c, req, 400)
		rr = httptest.NewRecorder()
	}
}

func (s *suite) Test_Register_success_with_minimal_metadata(c *C) {
	body := bytes.NewReader([]byte(`{"name": "my-app", "version": "1"}`))
	req, _ := http.NewRequest("POST", registerEndpoint, body)
	testRequest(c, req, 200)
}

func (s *suite) Test_GetApplications_empty_list(c *C) {
	req, _ := http.NewRequest("GET", applicationsEndpointNoProject, nil)
	testRequest(c, req, 404)
}

func (s *suite) Test_GetApplications(c *C) {
	s.addRelease(c, applicationsTestProject, "1")
	s.addRelease(c, applicationsTestProject, "2")
	req, _ := http.NewRequest("GET", applicationsEndpoint, nil)
	testRequest(c, req, http.StatusOK)
	c.Assert(rr.Body.String(), Equals, `["my-app"]`)
}

type hasItemChecker struct{}

var HasItem = &hasItemChecker{}

func (*hasItemChecker) Info() *CheckerInfo {
	return &CheckerInfo{Name: "HasItem", Params: []string{"obtained", "expected to have item"}}
}
func (*hasItemChecker) Check(params []interface{}, names []string) (bool, string) {
	obtained := params[0]
	expectedItem := params[1]
	switch obtained.(type) {
	case []interface{}:
		for _, v := range obtained.([]interface{}) {
			if v == expectedItem {
				return true, ""
			}
		}
	case []string:
		for _, v := range obtained.([]string) {
			if v == expectedItem {
				return true, ""
			}
		}
	default:
		return false, "Unexpected type."
	}
	return false, "Item not found"
}

func (s *suite) Test_GetVersions(c *C) {
	s.addRelease(c, applicationVersionsTestProject, "1")
	s.addRelease(c, applicationVersionsTestProject, "2")
	req, _ := http.NewRequest("GET", applicationVersionsEndpoint, nil)
	testRequest(c, req, http.StatusOK)
	result := []string{}
	err := json.Unmarshal([]byte(rr.Body.String()), &result)
	c.Assert(err, IsNil)
	c.Assert(result, HasLen, 2)
	c.Assert(result, HasItem, "1")
	c.Assert(result, HasItem, "2")
}

func (s *suite) Test_GetVersions_fails_if_app_not_found(c *C) {
	s.addRelease(c, applicationVersionsTestProject, "1")
	s.addRelease(c, applicationVersionsTestProject, "2")
	req, _ := http.NewRequest("GET", applicationVersionsNoApp, nil)
	testRequest(c, req, 404)
	rr = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", applicationVersionsNoProject, nil)
	testRequest(c, req, 404)
}
