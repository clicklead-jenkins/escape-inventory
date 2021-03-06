package sqlhelp

import (
	"database/sql"
	"encoding/json"
	"time"

	. "github.com/ankyra/escape-inventory/dao/types"
)

func (s *SQLHelper) AddApplication(app *Application) error {
	return s.PrepareAndExecInsert(s.AddApplicationQuery,
		app.Name,
		app.Project,
		app.Description,
		app.LatestVersion,
		app.Logo,
		app.UploadedBy,
		app.UploadedAt.Unix())
}
func (s *SQLHelper) UpdateApplication(app *Application) error {
	return s.PrepareAndExecUpdate(s.UpdateApplicationQuery,
		app.Description,
		app.LatestVersion,
		app.Logo,
		app.UploadedBy,
		app.UploadedAt.Unix(),
		app.Name,
		app.Project,
	)
}

func (s *SQLHelper) GetApplications(namespace string) (map[string]*Application, error) {
	rows, err := s.PrepareAndQuery(s.GetApplicationsQuery, namespace)
	if err != nil {
		return nil, err
	}
	return s.scanApplications(rows)
}

func (s *SQLHelper) GetApplication(namespace, name string) (*Application, error) {
	rows, err := s.PrepareAndQuery(s.GetApplicationQuery, namespace, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		return s.scanApplication(rows)
	}
	return nil, NotFound
}

func (s *SQLHelper) GetApplicationHooks(app *Application) (Hooks, error) {
	rows, err := s.PrepareAndQuery(s.GetApplicationHooksQuery, app.Project, app.Name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		return s.scanHooks(rows)
	}
	return nil, NotFound
}

func (s *SQLHelper) SetApplicationHooks(app *Application, hooks Hooks) error {
	bytes, err := json.Marshal(hooks)
	if err != nil {
		return err
	}
	return s.PrepareAndExecUpdate(s.SetApplicationHooksQuery,
		string(bytes),
		app.Project,
		app.Name)
}

func (s *SQLHelper) SetApplicationSubscribesToUpdatesFrom(app *Application, upstream []*Application) error {
	err := s.PrepareAndExec(s.DeleteSubscriptionsQuery,
		app.Project,
		app.Name)
	if err != nil {
		return err
	}
	for _, upstreamApp := range upstream {
		err := s.PrepareAndExecInsertIgnoreDups(s.AddSubscriptionQuery,
			app.Project,
			app.Name,
			upstreamApp.Project,
			upstreamApp.Name,
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *SQLHelper) GetDownstreamHooks(app *Application) ([]*Hooks, error) {
	rows, err := s.PrepareAndQuery(s.GetDownstreamSubscriptionsQuery, app.Project, app.Name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	result := []*Hooks{}
	for rows.Next() {
		var hooksString string
		if err := rows.Scan(&hooksString); err != nil {
			return nil, err
		}
		hooks := NewHooks()
		if err := json.Unmarshal([]byte(hooksString), &hooks); err != nil {
			return nil, err
		}
		result = append(result, &hooks)
	}
	return result, nil
}

func (s *SQLHelper) scanApplication(rows *sql.Rows) (*Application, error) {
	var name, namespace, description, latestVersion, logo, uploadedBy string
	var uploadedAt int64
	if err := rows.Scan(&name, &namespace, &description, &latestVersion, &logo, &uploadedBy, &uploadedAt); err != nil {
		return nil, err
	}
	return &Application{
		Name:          name,
		Project:       namespace,
		Description:   description,
		LatestVersion: latestVersion,
		Logo:          logo,
		UploadedBy:    uploadedBy,
		UploadedAt:    time.Unix(uploadedAt, 0),
	}, nil
}

func (s *SQLHelper) scanApplications(rows *sql.Rows) (map[string]*Application, error) {
	defer rows.Close()
	result := map[string]*Application{}
	for rows.Next() {
		app, err := s.scanApplication(rows)
		if err != nil {
			return nil, err
		}
		result[app.Name] = app
	}
	return result, nil
}
