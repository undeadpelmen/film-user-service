package main

import (
	"fmt"
	"os"

	"github.com/undeadpelmen/film-authorize/internal/logger"
	"github.com/undeadpelmen/film-authorize/internal/models"
)

func main() {

	if err := logger.Init(&logger.LoggerConfig{}); err != nil {
		fmt.Printf("Logger Init error: %v", err)
		os.Exit(1)
	}

	if err := models.Init(&models.DataBaseConfig{}); err != nil {
		fmt.Printf("DataBase Initialize error: %v", err)
		os.Exit(2)
	}
}
