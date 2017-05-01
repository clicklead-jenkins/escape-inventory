package main

import (
    "fmt"
    "os"
	"log"
	"net/http"
	"github.com/ankyra/escape-registry/config"
	"github.com/ankyra/escape-registry/handlers"
	"github.com/ankyra/escape-registry/dao"
	"github.com/ankyra/escape-registry/storage"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

const (
	defaultConfigFile = "config.json"
)

func loadConfig(configFile string) (*config.Config, error) {
    env := os.Environ()
    if !config.PathExists(configFile) {
        fmt.Println("Using default configuration")
        return config.NewConfig(env)
    } else {
        fmt.Printf("Loading configuration file '%s\n", configFile)
        return config.LoadConfig(configFile, env)
    }
}

func activateConfig(conf *config.Config) error {
    fmt.Printf("Activating '%s' database\n", conf.Database)
    if err := dao.LoadFromConfig(conf); err != nil {
        return err
    }
    fmt.Printf("Activating '%s' storage backend\n", conf.StorageBackend)
    if err := storage.LoadFromConfig(conf); err != nil {
        return err
    }
    return nil
}

func main() {
    conf, err := loadConfig(defaultConfigFile)
	if err != nil {
		panic(err)
	}
    if err := activateConfig(conf); err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.Handle("/r/", negroni.New(
		negroni.Wrap(http.HandlerFunc(handlers.RegisterHandler))))
	r.Handle("/r/{release}/", negroni.New(
		negroni.Wrap(http.HandlerFunc(handlers.GetMetadataHandler))))
	r.Handle("/r/{release}/download", negroni.New(
		negroni.Wrap(http.HandlerFunc(handlers.DownloadHandler))))
	r.Handle("/r/{release}/upload", negroni.New(
		negroni.Wrap(http.HandlerFunc(handlers.UploadHandler))))
	r.Handle("/r/{release}/next-version", negroni.New(
		negroni.Wrap(http.HandlerFunc(handlers.NextVersionHandler))))
	r.Handle("/export-releases", negroni.New(
		negroni.Wrap(http.HandlerFunc(handlers.ExportReleasesHandler))))
	r.Handle("/import-releases", negroni.New(
		negroni.Wrap(http.HandlerFunc(handlers.ImportReleasesHandler))))

	middleware := negroni.Classic()
	middleware.UseHandler(r)
	http.Handle("/", middleware)

	port := "3000"
	log.Printf("Starting Escape Registry v%s on port %s\n", registryVersion, port)
	log.Fatalln(http.ListenAndServe(":"+port, nil))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Escape Release Registry v" + registryVersion))
}
