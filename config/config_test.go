package config

import (
	"testing"
	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type configSuite struct{}

var _ = Suite(&configSuite{})

func (s *configSuite) Test_NewConfig_uses_Sqlite_and_local_storage_by_default(c *C) {
    env := []string{}
    conf, err := NewConfig(env)
    c.Assert(err, IsNil)
    c.Assert(conf.Database, Equals, "sqlite")
    c.Assert(conf.DatabaseSettings.Path, Equals, "/var/lib/escape/registry.db")
    c.Assert(conf.StorageBackend, Equals, "local")
    c.Assert(conf.StorageSettings.Path, Equals, "/var/lib/escape/releases")
}

func (s *configSuite) Test_LoadConfig_InMemoryDb(c *C) {
    env := []string{}
    conf, err := LoadConfig("testdata/in_memory_storage_config.json", env)
    c.Assert(err, IsNil)
    c.Assert(conf.Database, Equals, "memory")
}

func (s *configSuite) Test_LoadConfig_SqliteDb(c *C) {
    env := []string{}
    conf, err := LoadConfig("testdata/sqlite_storage_config.json", env)
    c.Assert(err, IsNil)
    c.Assert(conf.Database, Equals, "sqlite")
    c.Assert(conf.DatabaseSettings.Path, Equals, "/var/lib/escape/registry.db")
}

func (s *configSuite) Test_LoadConfig_LocalStorage(c *C) {
    env := []string{}
    conf, err := LoadConfig("testdata/local_storage_backend.json", env)
    c.Assert(err, IsNil)
    c.Assert(conf.StorageBackend, Equals, "local")
    c.Assert(conf.StorageSettings.Path, Equals, "/var/lib/escape/releases")
}

func (s *configSuite) Test_LoadConfig_GCS(c *C) {
    env := []string{}
    conf, err := LoadConfig("testdata/gcs_storage_backend.json", env)
    c.Assert(err, IsNil)
    c.Assert(conf.StorageBackend, Equals, "gcs")
    c.Assert(conf.StorageSettings.Path, Equals, "")
    c.Assert(conf.StorageSettings.Bucket, Equals, "gs://escape-releases/")
    c.Assert(conf.StorageSettings.Credentials["project-id"], Equals, "test")
}

func (s *configSuite) Test_LoadConfig_fails_if_not_exists(c *C) {
    env := []string{}
    conf, err := LoadConfig("testdata/doesnt_exist.json", env)
    c.Assert(conf, IsNil)
    c.Assert(err, Not(IsNil))
}

func (s *configSuite) Test_LoadConfig_fails_if_malformed(c *C) {
    env := []string{}
    conf, err := LoadConfig("testdata/malformed.json", env)
    c.Assert(conf, IsNil)
    c.Assert(err, Not(IsNil))
}


func (s *configSuite) Test_LoadConfig_Parses_Yaml(c *C) {
    env := []string{}
    conf, err := LoadConfig("testdata/yaml_config.yaml", env)
    c.Assert(err, IsNil)
    c.Assert(conf.Database, Equals, "sqlite")
    c.Assert(conf.DatabaseSettings.Path, Equals, "/var/lib/escape/registry.db")
    c.Assert(conf.StorageBackend, Equals, "gcs")
    c.Assert(conf.StorageSettings.Path, Equals, "")
    c.Assert(conf.StorageSettings.Bucket, Equals, "gs://escape-releases/")
    c.Assert(conf.StorageSettings.Credentials["project-id"], Equals, "test")
}

func (s *configSuite) Test_LoadConfig_Parses_Yml(c *C) {
    env := []string{}
    conf, err := LoadConfig("testdata/yml_config.yml", env)
    c.Assert(err, IsNil)
    c.Assert(conf.Database, Equals, "sqlite")
    c.Assert(conf.DatabaseSettings.Path, Equals, "/var/lib/escape/registry.db")
    c.Assert(conf.StorageBackend, Equals, "gcs")
    c.Assert(conf.StorageSettings.Path, Equals, "")
    c.Assert(conf.StorageSettings.Bucket, Equals, "gs://escape-releases/")
    c.Assert(conf.StorageSettings.Credentials["project-id"], Equals, "test")
}

func (s *configSuite) Test_LoadConfig_fails_if_yaml_malformed(c *C) {
    env := []string{}
    conf, err := LoadConfig("testdata/malformed.yaml", env)
    c.Assert(conf, IsNil)
    c.Assert(err, Not(IsNil))
}

func (s *configSuite) Test_NewConfig_Uses_EnvironmentVariables(c *C) {
    env := []string{
        "DATABASE=memory",
        "DATABASE_SETTINGS_PATH=",
        "STORAGE_BACKEND=gcs",
        "STORAGE_SETTINGS_PATH=",
        "STORAGE_SETTINGS_BUCKET=gs://escape-releases/",
        "STORAGE_SETTINGS_CREDENTIALS={\"project-id\": \"test\"}",
    }
    conf, err := NewConfig(env)
    c.Assert(err, IsNil)
    c.Assert(conf.Database, Equals, "memory")
    c.Assert(conf.DatabaseSettings.Path, Equals, "")
    c.Assert(conf.StorageBackend, Equals, "gcs")
    c.Assert(conf.StorageSettings.Path, Equals, "")
    c.Assert(conf.StorageSettings.Bucket, Equals, "gs://escape-releases/")
    c.Assert(conf.StorageSettings.Credentials["project-id"], Equals, "test")
}

func (s *configSuite) Test_NewConfig_Fails_If_Credentials_Malformed(c *C) {
    env := []string{
        "STORAGE_SETTINGS_CREDENTIALS={\"project-id\": \"t",
    }
    conf, err := NewConfig(env)
    c.Assert(conf, IsNil)
    c.Assert(err, Not(IsNil))
}

func (s *configSuite) Test_LoadConfig_Uses_EnvironmentVariables(c *C) {
    env := []string{
        "DATABASE=memory",
        "DATABASE_SETTINGS_PATH=",
        "STORAGE_BACKEND=local",
        "STORAGE_SETTINGS_PATH=/test/",
        "STORAGE_SETTINGS_BUCKET=",
        "STORAGE_SETTINGS_CREDENTIALS=",
    }
    conf, err := LoadConfig("testdata/yml_config.yml", env)
    c.Assert(err, IsNil)
    c.Assert(conf.Database, Equals, "memory")
    c.Assert(conf.DatabaseSettings.Path, Equals, "")
    c.Assert(conf.StorageBackend, Equals, "local")
    c.Assert(conf.StorageSettings.Path, Equals, "/test/")
    c.Assert(conf.StorageSettings.Bucket, Equals, "")
    c.Assert(conf.StorageSettings.Credentials["project-id"], Equals, "")
}