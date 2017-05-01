package handlers

import (
    "io/ioutil"
	"net/http"
    "encoding/json"
    "github.com/ankyra/escape-registry/model"
)

func ImportReleasesHandler(w http.ResponseWriter, r *http.Request) {
    releases, err := ioutil.ReadAll(r.Body)
    if err != nil {
        HandleError(w, r, err)
        return
    }
    releasesList := []map[string]interface{}{}
    if err := json.Unmarshal(releases, &releasesList); err != nil {
        HandleError(w, r, err)
        return
    }
    if err := model.Import(releasesList); err != nil {
        HandleError(w, r, err)
        return
    }
}