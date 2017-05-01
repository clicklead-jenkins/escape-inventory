package sqlite

import (
    sqlite3 "github.com/mattn/go-sqlite3"
    . "github.com/ankyra/escape-registry/dao/types"
)


type release_dao struct {
    dao *sql_dao
    releaseId string
    version string
    metadata Metadata
}

func newRelease(release Metadata, dao *sql_dao) *release_dao {
    return &release_dao{
        dao: dao,
        releaseId: release.GetReleaseId(),
        version: release.GetVersion(),
        metadata: release,
    }
}

func (r *release_dao) GetApplication() ApplicationDAO {
    return newApplicationDAO(
        r.metadata.GetType(),
        r.metadata.GetName(),
        r.dao,
    )
}

func (r *release_dao) GetVersion() string {
    return r.version
}

func (r *release_dao) GetMetadata() Metadata {
    return r.metadata
}

func (r *release_dao) GetPackageURIs() ([]string, error) {
    stmt, err := r.dao.db.Prepare("SELECT uri FROM package WHERE release_id = ?")
    if err != nil {
        return nil, err
    }
    rows, err := stmt.Query(r.releaseId)
    if err != nil {
        return nil, err
    }
    result := []string{}
    for rows.Next() {
        var uri string
        if err := rows.Scan(&uri); err != nil {
            return nil, err
        }
        result = append(result, uri)
    }
    return result, nil
}

func (r *release_dao) AddPackageURI(uri string) error {
    stmt, err := r.dao.db.Prepare("INSERT INTO package (release_id, uri) VALUES (?, ?)")
    if err != nil {
        return err
    }
    _, err = stmt.Exec(r.releaseId, uri)
    if err != nil {
        if err.(sqlite3.Error).Code == sqlite3.ErrConstraint {
            return AlreadyExists
        }
        return nil
    }
    return nil
}

func (r *release_dao) Save() (ReleaseDAO, error) {
    stmt, err := r.dao.db.Prepare(`
        INSERT INTO release(project, typ, name, release_id, version, metadata) VALUES(?, ?, ?, ?, ?, ?)`)
    if err != nil {
        return nil, err
    }
    project := ""
    typ := r.metadata.GetType()
    name := r.metadata.GetName()
    _, err = stmt.Exec(project, typ, name, r.releaseId, r.version, r.metadata.ToJson())
    if err != nil {
        if err.(sqlite3.Error).Code == sqlite3.ErrConstraint {
            return nil, AlreadyExists
        }
        return nil, err
    }
    return r, nil
}