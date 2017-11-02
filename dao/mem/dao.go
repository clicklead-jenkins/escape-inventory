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

package mem

import (
	. "github.com/ankyra/escape-inventory/dao/types"
)

type application struct {
	App      *Application
	Releases map[string]*release
}

type release struct {
	Release      *Release
	Packages     []string
	Dependencies []*Dependency
}

type dao struct {
	projectMetadata  map[string]*Project
	projectHooks     map[*Project]Hooks
	projects         map[string]map[string]*application
	apps             map[*Application]*application
	applicationHooks map[*Application]Hooks
	subscriptions    map[*Application][]*Application
	releases         map[*Release]*release
	acls             map[string]map[string]Permission
}

func NewInMemoryDAO() DAO {
	return &dao{
		projectMetadata:  map[string]*Project{},
		projectHooks:     map[*Project]Hooks{},
		projects:         map[string]map[string]*application{},
		apps:             map[*Application]*application{},
		applicationHooks: map[*Application]Hooks{},
		subscriptions:    map[*Application][]*Application{},
		releases:         map[*Release]*release{},
		acls:             map[string]map[string]Permission{},
	}
}
