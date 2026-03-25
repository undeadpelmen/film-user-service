package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dataBaseConnection *gorm.DB
	config             *DataBaseConfig
	isInit             bool
)

type DataBaseConfig struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
	sslmode  string
}

func initDb() error {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", config.host, config.port, config.user, config.password, config.dbname, config.sslmode)

	var err error

	if dataBaseConnection, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}); err != nil {
		//TODO: rewrite error handling
		return err
	}

	if err = dataBaseConnection.AutoMigrate(&User{}); err != nil {
		//TODO: rewrite error handling
		return err
	}

	return err
}

func Init(cfg *DataBaseConfig) error {
	config = &*cfg

	if isInit {
		//TODO: rewrite error handling
		return nil
	}

	if err := initDb(); err != nil {
		//TODO: rewrite error handling
		return err
	}

	return nil
}
