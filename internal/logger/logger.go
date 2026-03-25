package logger

import "github.com/rs/zerolog"

var (
	logger *zerolog.Logger
	config *LoggerConfig
	isInit bool
)

type LoggerConfig struct {
}

func initLogger() error {
	return nil
}

func Init(cfg *LoggerConfig) error {
	config = &*cfg

	if isInit {
		return nil
	}

	if err := initLogger(); err != nil {
		//TODO: rewrite error handling
		return err
	}

	return nil
}

func GetLogger() *zerolog.Logger {
	if isInit {
		return logger
	}

	return nil
}
