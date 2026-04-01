package logger

import (
	"io"
	"os"

	"github.com/rs/zerolog"
)

var (
	logger zerolog.Logger
	config *LoggerConfig
	isInit bool
)

type LoggerConfig struct {
	LogFilePath string
	LogLevel    zerolog.Level
}

func initLogger() error {
	var writer io.Writer

	if file, err := os.OpenFile("./log/"+config.LogFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777); err != nil {
		return err
	} else {
		writer = io.MultiWriter(file, zerolog.ConsoleWriter{Out: os.Stdout})
	}

	logger = zerolog.New(writer).With().Timestamp().Caller().Int("pid", os.Getpid()).Logger()

	return nil
}

func Init(cfg *LoggerConfig) error {
	config = cfg

	if isInit {
		return nil
	}

	if err := initLogger(); err != nil {
		//TODO: rewrite error handling
		return err
	}

	isInit = true

	return nil
}

func GetLogger() *zerolog.Logger {
	if isInit {
		return &logger
	}

	return nil
}
