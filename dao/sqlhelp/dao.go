package sqlhelp

import (
	"database/sql"
	"strconv"
	"strings"

	"github.com/ankyra/escape-core"
	. "github.com/ankyra/escape-registry/dao/types"
)

type SQLHelper struct {
	DB                       *sql.DB
	UseNumericInsertMarks    bool
	GetProjectQuery          string
	AddProjectQuery          string
	UpdateProjectQuery       string
	GetProjectsQuery         string
	GetProjectsByGroupsQuery string
	GetApplicationsQuery     string
	GetApplicationQuery      string
	FindAllVersionsQuery     string
	GetReleaseQuery          string
	GetAllReleasesQuery      string
	AddReleaseQuery          string
	GetPackageURIsQuery      string
	AddPackageURIQuery       string
	GetACLQuery              string
	InsertACLQuery           string
	UpdateACLQuery           string
	DeleteACLQuery           string
	GetPermittedGroupsQuery  string
	IsUniqueConstraintError  func(error) bool
}

func (s *SQLHelper) scanProject(rows *sql.Rows) (*Project, error) {
	var name, description, orgURL, logo string
	if err := rows.Scan(&name, &description, &orgURL, &logo); err != nil {
		return nil, err
	}
	return &Project{
		Name:        name,
		Description: description,
		OrgURL:      orgURL,
		Logo:        logo,
	}, nil
}

func (s *SQLHelper) scanProjects(rows *sql.Rows) (map[string]*Project, error) {
	defer rows.Close()
	result := map[string]*Project{}
	for rows.Next() {
		prj, err := s.scanProject(rows)
		if err != nil {
			return nil, err
		}
		result[prj.Name] = prj
	}
	return result, nil
}

func (s *SQLHelper) GetProject(project string) (*Project, error) {
	rows, err := s.PrepareAndQuery(s.GetProjectQuery, project)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		return s.scanProject(rows)
	}
	return nil, NotFound
}

func (s *SQLHelper) AddProject(project *Project) error {
	return s.PrepareAndExecInsert(s.AddProjectQuery,
		project.Name,
		project.Description,
		project.OrgURL,
		project.Logo)
}

func (s *SQLHelper) UpdateProject(project *Project) error {
	return s.PrepareAndExecUpdate(s.UpdateProjectQuery,
		project.Name,
		project.Description,
		project.OrgURL,
		project.Logo,
		project.Name)
}

func (s *SQLHelper) GetProjects() (map[string]*Project, error) {
	rows, err := s.PrepareAndQuery(s.GetProjectsQuery)
	if err != nil {
		return nil, err
	}
	return s.scanProjects(rows)
}

func (s *SQLHelper) GetProjectsByGroups(readGroups []string) (map[string]*Project, error) {
	starFound := false
	for _, g := range readGroups {
		if g == "*" {
			starFound = true
			break
		}
	}
	if !starFound {
		readGroups = append(readGroups, "*")
	}
	insertMarks := []string{}
	for i, _ := range readGroups {
		if s.UseNumericInsertMarks {
			insertMarks = append(insertMarks, "$"+strconv.Itoa(i+1))
		} else {
			insertMarks = append(insertMarks, "?")
		}
	}
	query := s.GetProjectsByGroupsQuery
	if len(readGroups) == 1 {
		query += " = " + insertMarks[0]
	} else {
		query += "IN (" + strings.Join(insertMarks, ", ") + ")"
	}
	interfaceGroups := []interface{}{}
	for _, g := range readGroups {
		interfaceGroups = append(interfaceGroups, g)
	}
	rows, err := s.PrepareAndQuery(query, interfaceGroups...)
	if err != nil {
		return nil, err
	}
	return s.scanProjects(rows)
}

func (s *SQLHelper) GetApplications(project string) ([]*Application, error) {
	rows, err := s.PrepareAndQuery(s.GetApplicationsQuery, project)
	if err != nil {
		return nil, err
	}
	apps, err := s.ReadRowsIntoStringArray(rows)
	if err != nil {
		return nil, err
	}
	result := []*Application{}
	for _, app := range apps {
		result = append(result, NewApplication(project, app))
	}
	return result, nil
}

func (s *SQLHelper) GetApplication(project, name string) (*Application, error) {
	rows, err := s.PrepareAndQuery(s.GetApplicationQuery, project, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		return NewApplication(project, name), nil
	}
	return nil, NotFound
}

func (s *SQLHelper) FindAllVersions(app *Application) ([]string, error) {
	rows, err := s.PrepareAndQuery(s.FindAllVersionsQuery, app.Project, app.Name)
	if err != nil {
		return nil, err
	}
	return s.ReadRowsIntoStringArray(rows)
}

func (s *SQLHelper) GetRelease(project, name, releaseId string) (*Release, error) {
	rows, err := s.PrepareAndQuery(s.GetReleaseQuery, project, name, releaseId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var metadataJson string
		if err := rows.Scan(&metadataJson); err != nil {
			return nil, err
		}
		metadata, err := core.NewReleaseMetadataFromJsonString(metadataJson)
		if err != nil {
			return nil, err
		}
		return NewRelease(NewApplication(project, name), metadata), nil
	}
	return nil, NotFound
}

func (s *SQLHelper) GetAllReleases() ([]*Release, error) {
	rows, err := s.PrepareAndQuery(s.GetAllReleasesQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	result := []*Release{}
	for rows.Next() {
		var project, metadataJson string
		if err := rows.Scan(&project, &metadataJson); err != nil {
			return nil, err
		}
		metadata, err := core.NewReleaseMetadataFromJsonString(metadataJson)
		if err != nil {
			return nil, err
		}
		result = append(result, NewRelease(NewApplication(project, metadata.Name), metadata))
	}
	return result, nil
}

func (s *SQLHelper) AddRelease(project string, release *core.ReleaseMetadata) (*Release, error) {
	stmt, err := s.DB.Prepare(s.AddReleaseQuery)
	if err != nil {
		return nil, err
	}
	name := release.Name
	_, err = stmt.Exec(project, name, release.GetReleaseId(), release.Version, release.ToJson())
	if err != nil {
		if s.IsUniqueConstraintError(err) {
			return nil, AlreadyExists
		}
		return nil, err
	}
	return NewRelease(NewApplication(project, release.Name), release), nil
}

func (s *SQLHelper) GetPackageURIs(release *Release) ([]string, error) {
	rows, err := s.PrepareAndQuery(s.GetPackageURIsQuery, release.Application.Project, release.ReleaseId)
	if err != nil {
		return nil, err
	}
	return s.ReadRowsIntoStringArray(rows)
}

func (s *SQLHelper) AddPackageURI(release *Release, uri string) error {
	stmt, err := s.DB.Prepare(s.AddPackageURIQuery)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(release.Application.Project, release.ReleaseId, uri)
	if err != nil {
		if s.IsUniqueConstraintError(err) {
			return AlreadyExists
		}
		return err
	}
	return nil
}

func (s *SQLHelper) SetACL(project, group string, perm Permission) error {
	stmt, err := s.DB.Prepare(s.InsertACLQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(project, group, int(perm))
	if err != nil {
		if s.IsUniqueConstraintError(err) {
			stmt, err := s.DB.Prepare(s.UpdateACLQuery)
			if err != nil {
				return err
			}
			_, err = stmt.Exec(int(perm), project, group)
			return err
		}
		return err
	}
	return nil
}

func (s *SQLHelper) GetACL(project string) (map[string]Permission, error) {
	rows, err := s.PrepareAndQuery(s.GetACLQuery, project)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	result := map[string]Permission{}
	for rows.Next() {
		var group_name string
		var permission Permission
		if err := rows.Scan(&group_name, &permission); err != nil {
			return nil, err
		}
		result[group_name] = permission
	}
	return result, nil
}

func (s *SQLHelper) DeleteACL(project, group string) error {
	stmt, err := s.DB.Prepare(s.DeleteACLQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(project, group)
	return err
}

func (s *SQLHelper) GetPermittedGroups(project string, perm Permission) ([]string, error) {
	rows, err := s.PrepareAndQuery(s.GetPermittedGroupsQuery, project, int(perm))
	if err != nil {
		return nil, err
	}
	return s.ReadRowsIntoStringArray(rows)
}

func (s *SQLHelper) ReadRowsIntoStringArray(rows *sql.Rows) ([]string, error) {
	defer rows.Close()
	result := []string{}
	for rows.Next() {
		var arg string
		if err := rows.Scan(&arg); err != nil {
			return nil, err
		}
		result = append(result, arg)
	}
	return result, nil
}

func (s *SQLHelper) PrepareAndQuery(query string, arg ...interface{}) (*sql.Rows, error) {
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Query(arg...)
}

func (s *SQLHelper) PrepareAndExecInsert(query string, arg ...interface{}) error {
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(arg...)
	if err != nil {
		if s.IsUniqueConstraintError(err) {
			return AlreadyExists
		}
	}
	return err
}

func (s *SQLHelper) PrepareAndExecUpdate(query string, arg ...interface{}) error {
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(arg...)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return NotFound
	}
	return err
}
