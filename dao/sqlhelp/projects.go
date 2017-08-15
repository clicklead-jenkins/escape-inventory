package sqlhelp

import (
	"database/sql"
	"strconv"
	"strings"

	. "github.com/ankyra/escape-registry/dao/types"
)

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