// Code generated by go-bindata.
// sources:
// dao/ql/schemas/1_initial_schema.down.sql
// dao/ql/schemas/1_initial_schema.up.sql
// dao/ql/schemas/2_metrics.down.sql
// dao/ql/schemas/2_metrics.up.sql
// dao/ql/schemas/3_metrics_user_id.down.sql
// dao/ql/schemas/3_metrics_user_id.up.sql
// dao/ql/schemas/4_feeds.up.sql
// dao/ql/schemas/5_feeds_application_field.up.sql
// dao/ql/schemas/6_providers.up.sql
// dao/ql/schemas/7_remove_feeds.sql
// DO NOT EDIT!

package ql

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var __1_initial_schemaDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\x28\x4a\xcd\x49\x4d\x2c\x4e\xb5\xe6\x42\x12\x2b\x48\x4c\xce\x4e\x4c\x47\x15\x4b\x4c\xce\x41\x55\x53\x94\x9f\x95\x9a\x5c\x82\xaa\xa6\xa0\x20\x27\x33\x39\xb1\x24\x33\x3f\x0f\x45\x1c\x6a\x47\x7c\x4a\x6a\x41\x6a\x5e\x4a\x6a\x5e\x72\x25\x8a\x74\x71\x69\x52\x71\x72\x51\x66\x01\x48\x5f\xb1\x35\x20\x00\x00\xff\xff\xb3\x3e\xc0\xc0\x9c\x00\x00\x00")

func _1_initial_schemaDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__1_initial_schemaDownSql,
		"1_initial_schema.down.sql",
	)
}

func _1_initial_schemaDownSql() (*asset, error) {
	bytes, err := _1_initial_schemaDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1_initial_schema.down.sql", size: 156, mode: os.FileMode(420), modTime: time.Unix(1511516577, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __1_initial_schemaUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x94\x41\x8f\xd3\x30\x10\x85\xcf\xe9\xaf\x18\xf5\xb4\x48\x41\x82\x33\xa7\xc2\x66\xa5\x4a\x55\x57\xec\xa6\xd2\xde\x22\x27\x9e\x0d\xa6\xae\x6d\xd9\x0e\x50\x10\xff\x1d\x35\x89\xdd\x38\x4e\xb7\x01\xa1\xbd\xbe\x4c\xec\x99\x6f\x9e\xdf\xe2\xd3\x43\xb6\xca\x33\xc8\x57\x1f\x37\x19\xac\xef\x60\x7b\x9f\x43\xf6\xb4\x7e\xcc\x1f\x41\x23\x47\x62\x10\x6e\x16\x00\x00\x82\x1c\x10\x8c\xd5\x4c\xd4\x29\xb4\x4a\xff\xbd\x60\xd4\xe9\xad\xfc\x0d\xb5\x61\x52\x04\xda\x01\x2d\xa1\xc4\x12\x28\xb9\x2c\x3b\x49\x69\xf9\x15\x2b\x1b\x94\x29\x2d\x2b\x34\x06\x69\x41\x51\xa1\xa0\x28\x2a\x86\x06\x4a\x29\x39\xdc\x66\x77\xab\xdd\x26\x87\x67\xc2\x0d\x76\xe5\x54\x7e\x17\x5c\x12\x6a\x80\x09\xeb\x0b\xde\x75\x1f\x1b\x75\xfa\x84\xb4\x28\x8f\xfd\x1d\xbe\x62\xb9\x1c\x95\x10\x3b\x3e\xe1\xcd\x87\x85\x43\xb3\xdb\xae\x3f\xef\x32\x58\x6f\x6f\xb3\xa7\x69\x42\x85\xda\xc3\xfd\xf6\xcc\xeb\x84\x2a\x75\x1c\x52\x37\xe9\xe0\xc8\x29\xda\x8a\x54\x7b\x52\x3b\xda\x53\x74\x62\xde\xdd\x1e\x1a\xcd\x42\xe1\x99\x71\x34\xec\x27\x06\x43\xbd\x7d\xff\x7a\x5c\xfa\x59\x7a\x2e\x7e\xb2\xf3\x00\xe9\xa9\xe9\xb9\x64\x48\xc5\x5f\xa0\x52\x6b\xd9\xa8\x22\x76\xa7\x42\x7d\x60\xa6\x75\x22\x13\x76\x6e\xe7\xa4\xe2\x7d\xd7\xed\xad\xfd\x85\xe9\xe0\x96\x6b\x6b\xec\x5b\x8c\x1f\x4d\x67\x59\x34\x95\x66\xca\x8e\x1f\x88\xd4\xf5\xee\x61\x13\x48\x5c\xd6\x32\x10\xbe\x48\xb9\x37\xd1\xce\x7e\xfd\x5e\xce\x5e\x4b\xd7\x9b\x5b\x8b\xeb\x74\xc6\x54\x44\x29\xce\x2a\xd2\xb6\x7d\x29\x0e\xa6\x96\x13\x8f\x1b\x99\x8d\x13\x8b\xc6\x16\x61\x6a\xc4\x55\x67\x18\x97\xed\xfa\x4f\x8e\xfe\x2f\x64\x07\x7c\x9c\x7d\x86\xc4\xba\x40\x98\xe7\x76\xf7\x48\x7c\x02\x1e\x5f\x30\x7f\xe4\xaf\xa9\xf0\xa5\xa8\x8a\xe9\xe5\x84\xef\xc6\x8b\x53\x87\x94\x0d\xe3\xb4\x30\x95\x54\xd8\xe6\xb1\xaf\xe6\xf2\x18\xc9\xcc\x14\xf8\xc3\xa2\x68\x4f\x99\x4a\xef\xbf\xcc\xd7\x33\x8b\x30\x6a\x03\x48\xfe\xad\x8e\xe2\x77\x30\x7e\xea\x67\x4e\x87\x83\x5e\xd9\x88\x69\x4a\xef\x62\x03\x37\x8b\x24\x80\x96\x8c\xd1\x26\xc3\xfa\x10\x70\xf8\x69\xfc\xe3\x25\x26\xc1\xfd\x3d\x80\x51\x4f\x81\xc1\x52\x88\x3a\x18\x49\xde\x89\x49\x02\x7f\x02\x00\x00\xff\xff\xe6\xbd\xee\x40\x02\x08\x00\x00")

func _1_initial_schemaUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1_initial_schemaUpSql,
		"1_initial_schema.up.sql",
	)
}

func _1_initial_schemaUpSql() (*asset, error) {
	bytes, err := _1_initial_schemaUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1_initial_schema.up.sql", size: 2050, mode: os.FileMode(420), modTime: time.Unix(1511516577, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __2_metricsDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xc8\x4d\x2d\x29\xca\x4c\x2e\xb6\xe6\x02\x8b\x79\xfa\xb9\xb8\x46\xc0\xc4\xe2\x0b\xb2\xad\xb9\x00\x01\x00\x00\xff\xff\x18\x02\x28\x6c\x2b\x00\x00\x00")

func _2_metricsDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__2_metricsDownSql,
		"2_metrics.down.sql",
	)
}

func _2_metricsDownSql() (*asset, error) {
	bytes, err := _2_metricsDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "2_metrics.down.sql", size: 43, mode: os.FileMode(420), modTime: time.Unix(1515079328, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __2_metricsUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\xcd\xb1\x0a\x83\x30\x14\x46\xe1\x3d\x4f\xf1\x8f\x15\x1c\xba\x3b\xa5\xf5\x0a\x01\x89\xb4\xde\x80\x9b\x94\x10\x4a\x5a\x8c\x92\x5c\xdf\xbf\x50\x2a\x5d\xcf\x70\xbe\xeb\x9d\x34\x13\x58\x5f\x7a\xc2\x12\x24\x47\x5f\x70\x52\x00\xb0\x97\x90\xd3\x63\x09\x28\x92\x63\x7a\xd6\xdf\xb8\xe5\xf5\x15\xbc\xcc\x7e\xdd\x93\x20\x26\x41\x4b\x9d\x76\x3d\xe3\x5c\xab\xaa\x51\xea\x37\x74\xd6\xdc\x1c\xc1\xd8\x96\x26\x98\x0e\x76\x60\xd0\x64\x46\x1e\x0f\x65\xde\xde\x18\xec\xdf\x3c\xb8\xaa\x51\x9f\x00\x00\x00\xff\xff\x1c\x82\x9d\xbd\x95\x00\x00\x00")

func _2_metricsUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__2_metricsUpSql,
		"2_metrics.up.sql",
	)
}

func _2_metricsUpSql() (*asset, error) {
	bytes, err := _2_metricsUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "2_metrics.up.sql", size: 149, mode: os.FileMode(420), modTime: time.Unix(1515079328, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __3_metrics_user_idDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\xce\xb1\x0a\x83\x30\x14\x85\xe1\xfd\x3e\xc5\x19\x2b\x38\x74\x77\xb2\xf5\x0a\x01\x89\xad\x26\xe0\x26\x25\x84\x92\x16\xa3\x24\xf1\xfd\x0b\x52\xb1\x5d\xbf\xe1\x9c\xbf\xea\xda\x1b\x84\xac\x78\xc0\x64\x53\x70\x26\x8e\xcb\xbb\xa0\x8d\x55\x79\x69\x78\xe7\x82\xe8\xda\x71\xa9\xf8\x9f\x71\x22\x00\x58\xa3\x0d\xfe\x31\x59\xc4\x14\x9c\x7f\xe6\x1b\x2e\x61\x7e\x59\x93\x46\x33\xaf\x3e\xc1\xf9\x84\x8a\xeb\x52\x37\x0a\xe7\x9c\xb2\x63\x50\x4b\x71\xd7\xfc\xad\x10\x35\x64\xab\xc0\x83\xe8\x55\xff\xd3\x84\x56\x1e\x9f\xfb\x5d\x56\xd0\x27\x00\x00\xff\xff\x9c\x4b\xd0\xe1\xc1\x00\x00\x00")

func _3_metrics_user_idDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__3_metrics_user_idDownSql,
		"3_metrics_user_id.down.sql",
	)
}

func _3_metrics_user_idDownSql() (*asset, error) {
	bytes, err := _3_metrics_user_idDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "3_metrics_user_id.down.sql", size: 193, mode: os.FileMode(420), modTime: time.Unix(1515081776, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __3_metrics_user_idUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\xce\xb1\xaa\x83\x30\x18\xc5\xf1\xfd\x7b\x8a\x33\x5e\xc1\xe1\xee\x4e\xb6\x7e\x42\x40\x62\xab\x09\xb8\x65\x48\x43\x49\x4b\xa3\x24\xf1\xfd\x0b\x62\xb1\x5d\x7f\xc3\x39\xff\x66\xe8\x2f\x10\xb2\xe1\x09\x2f\x97\xa3\xb7\xc9\x2c\xcf\x8a\x36\x56\xf5\xa9\xe3\x0f\x57\x44\xe7\x81\x6b\xc5\xbf\x8c\x3f\x02\x80\x35\xb9\x68\xfc\x0d\x29\x47\x1f\xee\xe5\x66\x4b\x9c\x1f\xce\x66\x63\xe7\x35\x64\xf8\x90\xd1\x70\x5b\xeb\x4e\xe1\xbf\xa4\xe2\xd8\xd3\x52\x5c\x35\xef\x11\xa2\x85\xec\x15\x78\x12\xa3\x1a\xbf\x92\xd0\xcb\xe3\x72\x7f\x2b\x2a\x7a\x07\x00\x00\xff\xff\xc3\x37\x4c\x96\xbf\x00\x00\x00")

func _3_metrics_user_idUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__3_metrics_user_idUpSql,
		"3_metrics_user_id.up.sql",
	)
}

func _3_metrics_user_idUpSql() (*asset, error) {
	bytes, err := _3_metrics_user_idUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "3_metrics_user_id.up.sql", size: 191, mode: os.FileMode(420), modTime: time.Unix(1515081776, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __4_feedsUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x8e\x41\xcb\x82\x40\x10\x86\xef\xfb\x2b\x5e\x3c\x29\x7c\x87\xef\xee\xc9\x72\x04\x41\x14\x72\x03\x6f\xb2\xe4\x14\x06\x6e\xb2\x3b\x05\xfd\xfb\x40\xb7\xf0\x50\xc7\x79\x1e\x86\xf7\xd9\x1f\x28\xd3\x04\x9d\xed\x2a\xc2\x99\x79\xe8\xf9\xc1\x56\x3c\x62\x05\x00\xcb\xd1\xcb\x73\x66\x78\x71\xa3\xbd\xfc\x2d\xf8\xee\xd9\x59\x33\xbd\x21\x72\x2a\xb2\x63\xa5\x11\x45\xab\x9f\xdd\xed\xca\x27\xf9\xa5\x65\x9c\xd8\x8b\x99\x66\x8c\x56\x3e\xf6\x7f\x95\x83\x11\xf3\xed\x31\x49\x95\x0a\xb5\x65\x9d\x53\x87\xb2\x40\xdd\x68\x50\x57\xb6\xba\xdd\xb6\xf7\x61\xde\xa3\xa9\xb7\x3c\x0e\x3c\x49\xd5\x2b\x00\x00\xff\xff\x2a\xca\xc3\xbb\xf7\x00\x00\x00")

func _4_feedsUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__4_feedsUpSql,
		"4_feeds.up.sql",
	)
}

func _4_feedsUpSql() (*asset, error) {
	bytes, err := _4_feedsUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "4_feeds.up.sql", size: 247, mode: os.FileMode(420), modTime: time.Unix(1515506470, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __5_feeds_application_fieldUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x93\x41\x6b\x84\x30\x10\x85\xef\xf9\x15\x8f\x3d\xad\xe0\xa1\x77\x4f\xae\xc6\x12\xd0\xa4\x68\x0a\x7b\x93\xb0\x9b\x16\x4b\x75\x83\xa6\x85\xfe\xfb\x42\x56\x57\x5d\xb4\xc5\x1e\xeb\xcd\xcc\xcb\xcc\xbc\x7c\x33\x07\xfa\xc8\x38\x64\x1e\xf2\x22\x8c\x24\x13\x3c\x20\x40\x9c\x8b\x27\x30\x1e\xd3\x23\x5e\xb4\x3e\x97\xfa\x53\x37\xb6\x2b\x4d\x7b\x79\xd3\x27\xdb\x05\x84\x00\x51\x4e\x43\x49\x21\xc3\x43\x4a\x61\x6b\x53\x4e\x94\xd8\x13\xb8\xcf\xfd\x96\xf6\xcb\x68\x74\xb6\xad\x9a\x57\xbf\x0f\x7c\x74\xba\x6d\x54\x3d\x1c\x23\xa6\x49\xf8\x9c\x4a\xec\x76\x83\xa2\x2f\xb6\x2e\x50\xc6\xbc\x57\x27\x65\xab\x4b\xb3\x2e\xb2\x55\xad\x3b\xab\x6a\x83\xaa\xb1\xb7\xf8\xc3\x10\x3e\x2b\xab\x96\x2f\x7b\xce\x24\xe3\x05\xcd\x25\x18\x97\xe2\xde\xe3\x7e\xf4\xe6\xdf\xec\xf8\x43\xdb\xfe\x58\xd9\x77\x55\x3c\x14\x34\xa5\x91\xc4\xb6\x6b\x48\x72\x91\x4d\x21\xb8\xb6\x1c\x9f\xeb\xcb\xcf\x42\x91\xc8\x32\x26\x03\x42\x16\xa9\xce\x88\xfd\x77\x5a\x1b\x48\x4d\x5a\xfb\x33\xb6\x9f\x72\xfc\xca\xf0\x6e\xb2\xa6\xeb\x75\x5d\x42\x96\x80\x0b\x09\x7a\x64\x85\x2c\x16\x57\x12\x82\xcf\x2c\xf7\xe7\x5e\xb0\x39\x55\x39\xb1\xb2\x96\x76\xe6\xd7\x1b\xe7\xee\x3b\x00\x00\xff\xff\xbf\x75\x58\x81\x4d\x04\x00\x00")

func _5_feeds_application_fieldUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__5_feeds_application_fieldUpSql,
		"5_feeds_application_field.up.sql",
	)
}

func _5_feeds_application_fieldUpSql() (*asset, error) {
	bytes, err := _5_feeds_application_fieldUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "5_feeds_application_field.up.sql", size: 1101, mode: os.FileMode(420), modTime: time.Unix(1519986076, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __6_providersUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x8e\x41\x0a\xc2\x30\x10\x45\xf7\x39\xc5\x5f\xb6\x90\x1b\x74\x55\x75\x84\x80\xa4\x68\x53\xe8\x4e\x4a\x1a\x24\x2a\x6d\x48\x42\xcf\x2f\x86\xb4\x16\x5c\xce\xfb\x33\xf3\xdf\xf1\x46\xb5\x22\xa8\xfa\x70\x21\x38\x3f\x2f\x76\x34\x3e\xa0\x60\x00\xbe\xf3\xd3\xe8\x88\x10\xbd\x9d\x1e\x3c\xb1\xc1\xb9\xb7\xd5\x43\xb4\xf3\xb4\x72\xa4\x60\x31\x3e\xfc\xc1\xd1\x04\xed\xad\xdb\x6f\xaf\x9f\x53\x53\x86\xac\xac\x18\xcb\x2a\x9d\x14\xd7\x8e\x20\xe4\x89\x7a\x88\x33\x64\xa3\x40\xbd\x68\x55\xbb\x5d\xdd\xdd\x0b\x8d\xfc\xe9\x16\x59\x94\xef\xed\xf8\x96\x97\x15\xfb\x04\x00\x00\xff\xff\x11\xc7\xeb\x0b\xe7\x00\x00\x00")

func _6_providersUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__6_providersUpSql,
		"6_providers.up.sql",
	)
}

func _6_providersUpSql() (*asset, error) {
	bytes, err := _6_providersUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "6_providers.up.sql", size: 231, mode: os.FileMode(420), modTime: time.Unix(1519986076, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __7_remove_feedsSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\x48\x4b\x4d\x4d\x89\x4f\x2d\x4b\xcd\x2b\x29\xb6\x06\x04\x00\x00\xff\xff\xd3\xfe\xd7\x50\x17\x00\x00\x00")

func _7_remove_feedsSqlBytes() ([]byte, error) {
	return bindataRead(
		__7_remove_feedsSql,
		"7_remove_feeds.sql",
	)
}

func _7_remove_feedsSql() (*asset, error) {
	bytes, err := _7_remove_feedsSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "7_remove_feeds.sql", size: 23, mode: os.FileMode(420), modTime: time.Unix(1523030428, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"1_initial_schema.down.sql": _1_initial_schemaDownSql,
	"1_initial_schema.up.sql": _1_initial_schemaUpSql,
	"2_metrics.down.sql": _2_metricsDownSql,
	"2_metrics.up.sql": _2_metricsUpSql,
	"3_metrics_user_id.down.sql": _3_metrics_user_idDownSql,
	"3_metrics_user_id.up.sql": _3_metrics_user_idUpSql,
	"4_feeds.up.sql": _4_feedsUpSql,
	"5_feeds_application_field.up.sql": _5_feeds_application_fieldUpSql,
	"6_providers.up.sql": _6_providersUpSql,
	"7_remove_feeds.sql": _7_remove_feedsSql,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"1_initial_schema.down.sql": &bintree{_1_initial_schemaDownSql, map[string]*bintree{}},
	"1_initial_schema.up.sql": &bintree{_1_initial_schemaUpSql, map[string]*bintree{}},
	"2_metrics.down.sql": &bintree{_2_metricsDownSql, map[string]*bintree{}},
	"2_metrics.up.sql": &bintree{_2_metricsUpSql, map[string]*bintree{}},
	"3_metrics_user_id.down.sql": &bintree{_3_metrics_user_idDownSql, map[string]*bintree{}},
	"3_metrics_user_id.up.sql": &bintree{_3_metrics_user_idUpSql, map[string]*bintree{}},
	"4_feeds.up.sql": &bintree{_4_feedsUpSql, map[string]*bintree{}},
	"5_feeds_application_field.up.sql": &bintree{_5_feeds_application_fieldUpSql, map[string]*bintree{}},
	"6_providers.up.sql": &bintree{_6_providersUpSql, map[string]*bintree{}},
	"7_remove_feeds.sql": &bintree{_7_remove_feedsSql, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

