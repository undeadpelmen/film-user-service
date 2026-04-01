package models

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	dataBaseConnection *gorm.DB
	config             *DataBaseConfig
	isInit             bool
)

type DataBaseConfig struct {
	Host     string
	Port     uint
	User     string
	Password string
	Dbname   string
	Sslmode  string
	logFile  string
}

func initDb() error {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", config.Host, config.Port, config.User, config.Password, config.Dbname, config.Sslmode)

	var err error

	if dataBaseConnection, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  logger.Info,
				IgnoreRecordNotFoundError: false,
				ParameterizedQueries:      false,
				Colorful:                  true,
			},
		),
	}); err != nil {
		//TODO: rewrite error handling
		return err
	}

	//create user table
	if err = dataBaseConnection.AutoMigrate(&User{}); err != nil {
		//TODO: rewrite error handling
		return err
	}

	////create settings table
	//if err = dataBaseConnection.AutoMigrate(&Settins{}); err != nil {
	//	// TODO: reewite error handling
	//	return err
	//}

	return err
}

func Init(cfg *DataBaseConfig) error {
	config = cfg

	if isInit {
		//TODO: rewrite error handling
		return nil
	}

	if err := initDb(); err != nil {
		//TODO: rewrite error handling
		return err
	}

	isInit = true

	return nil
}
