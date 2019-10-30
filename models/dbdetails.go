package models

import (
	"github.com/go-yaml/yaml"
	"io/ioutil"
)

type DbAccess struct {
	DbUser string `yaml: "dbuser"`
	DbPass string `yaml: "dbpass"`
	DbType string `yaml: "dbtype"`
	DbName string `yaml: "dbname"`
	DbHost string `yaml: "dbhost"`
	DbPort int    `yaml: "dbport"`
	DbPath string `yaml: "dbpath"`
}

func (db *DbAccess) Init(configPath string) (err error) {

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
