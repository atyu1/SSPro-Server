package datapoints

import (
	"io/ioutil"
	"github.com/go-yaml/yaml"
)

type DbAccess struct {
	Username string `yaml: dbuser`
	Password string `yaml: dbpass`
	Type	string  `yaml: dbtype`
	DbName	string  `yaml: dbname`
	DbHost	string  `yaml: dbhost`
	DbPort	int     `yaml: dbport`
	DbPath	string	`yaml: sqlite3_path`
}

func (db DbAccess) Init(configPath string) (err error) {
	
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, &db)
	if err != nil {
		return err
	}
	return nil
}
