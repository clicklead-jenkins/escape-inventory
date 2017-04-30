package types

import (
    "github.com/ankyra/escape-client/model/interfaces"
    "fmt"
)

type Metadata interfaces.ReleaseMetadata

type DAO interface {
    GetApplication(typ, name string) (ApplicationDAO, error)
    GetApplications() ([]ApplicationDAO, error)
    GetRelease(releaseId string) (ReleaseDAO, error)
    GetAllReleases() ([]ReleaseDAO, error)
    AddRelease(metadata Metadata) (ReleaseDAO, error)
}

type ApplicationDAO interface {
    GetType() string
    GetName() string
    FindAllVersions() ([]string, error)
}

type ReleaseDAO interface {
    GetApplication() ApplicationDAO
    GetVersion() string
    GetMetadata() Metadata

    GetPackageURIs() ([]string, error)
    AddPackageURI(uri string) error
}

var NotFound = fmt.Errorf("Not found")
var AlreadyExists = fmt.Errorf("Already exists")
