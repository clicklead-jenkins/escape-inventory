/*
Copyright 2017, 2018 Ankyra

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
	"log"
	"net/http"

	"github.com/ankyra/escape-inventory/cmd"
	"github.com/ankyra/escape-inventory/config"
	"github.com/ankyra/escape-inventory/handlers"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var ReadRoutes = map[string]http.HandlerFunc{
	"/":       HomeHandler,
	"/health": handlers.HealthCheckHandler,

	"/api/v1/inventory/":                                                             handlers.GetNamespacesHandler,
	"/api/v1/inventory/{namespace}/":                                                 handlers.GetNamespaceHandler,
	"/api/v1/inventory/{namespace}/hooks/":                                           handlers.GetNamespaceHooksHandler,
	"/api/v1/inventory/{namespace}/units/":                                           handlers.GetApplicationsHandler,
	"/api/v1/inventory/{namespace}/units/{name}/":                                    handlers.GetApplicationHandler,
	"/api/v1/inventory/{namespace}/units/{name}/hooks/":                              handlers.GetApplicationHooksHandler,
	"/api/v1/inventory/{namespace}/units/{name}/versions/":                           handlers.GetApplicationVersionsHandler,
	"/api/v1/inventory/{namespace}/units/{name}/versions/{version}/":                 handlers.GetVersionHandler,
	"/api/v1/inventory/{namespace}/units/{name}/versions/{version}/downstream":       handlers.DownstreamHandler,
	"/api/v1/inventory/{namespace}/units/{name}/versions/{version}/dependency-graph": handlers.DependencyGraphHandler,
	"/api/v1/inventory/{namespace}/units/{name}/versions/{version}/diff/":            handlers.DiffHandler,
	"/api/v1/inventory/{namespace}/units/{name}/versions/{version}/diff/{diffWith}/": handlers.DiffHandler,
	"/api/v1/inventory/{namespace}/units/{name}/versions/{version}/download":         handlers.DownloadHandler,
	"/api/v1/inventory/{namespace}/units/{name}/versions/{version}/previous/":        handlers.PreviousVersionHandler,
	"/api/v1/inventory/{namespace}/units/{name}/next-version":                        handlers.NextVersionHandler,
	"/api/v1/inventory/__providers":                                                  handlers.ProviderHandler,
}

var DeleteRoutes = map[string]http.HandlerFunc{
	"/api/v1/inventory/{namespace}/hard-delete": handlers.HardDeleteNamespaceHandler,
}

var WriteRoutes = map[string]http.HandlerFunc{
	"/api/v1/inventory/{namespace}/add-namespace":                          handlers.AddNamespaceHandler,
	"/api/v1/inventory/{namespace}/register":                               handlers.RegisterHandler,
	"/api/v1/inventory/{namespace}/units/{name}/tags/":                     handlers.TagReleaseHandler,
	"/api/v1/inventory/{namespace}/units/{name}/versions/{version}/upload": handlers.UploadHandler,
}

var UpdateRoutes = map[string]http.HandlerFunc{
	"/api/v1/inventory/{namespace}/":                    handlers.UpdateNamespaceHandler,
	"/api/v1/inventory/{namespace}/hooks/":              handlers.UpdateNamespaceHooksHandler,
	"/api/v1/inventory/{namespace}/units/{name}/hooks/": handlers.UpdateApplicationHooksHandler,
}

var DevRoutes = map[string]map[string]http.HandlerFunc{
	"/api/v1/internal/database": {
		"DELETE": handlers.WipeDatabaseHandler,
	},
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Escape Release Inventory v" + cmd.InventoryVersion))
}

func getMux(cfg *config.Config) *mux.Router {
	r := mux.NewRouter()
	getRouter := r.Methods("GET").Subrouter()
	for url, handler := range ReadRoutes {
		getRouter.Handle(url, handler)
	}
	postRouter := r.Methods("POST").Subrouter()
	for url, handler := range WriteRoutes {
		postRouter.Handle(url, handler)
	}
	putRouter := r.Methods("PUT").Subrouter()
	for url, handler := range UpdateRoutes {
		putRouter.Handle(url, handler)
	}
	deleteRouter := r.Methods("DELETE").Subrouter()
	for url, handler := range DeleteRoutes {
		deleteRouter.Handle(url, handler)
	}

	if cfg.Dev {
		for url, methodHandlers := range DevRoutes {
			for method, handler := range methodHandlers {
				r.Methods(method).Subrouter().Handle(url, handler)
			}
		}
	}

	r.Handle("/metrics", promhttp.Handler())
	return r
}

func main() {
	cfg := cmd.LoadConfig()
	if cfg.Dev {
		log.Println("INFO: Starting in Dev mode.")
	}
	cmd.StartInventory(getMux(cfg))
}
