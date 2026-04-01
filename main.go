package main

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/undeadpelmen/film-authorize/internal/logger"
	"github.com/undeadpelmen/film-authorize/internal/models"
)

var (
	lg *zerolog.Logger
)

func main() {

	if err := logger.Init(&logger.LoggerConfig{
		LogLevel:    zerolog.TraceLevel,
		LogFilePath: "auth.log",
	}); err != nil {
		fmt.Printf("Logger Init error: %v", err)
		os.Exit(1)
	}

	lg = logger.GetLogger()
	if lg == nil {
		fmt.Println("Logger iz nil")
		os.Exit(1)
	}

	if err := models.Init(&models.DataBaseConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "mylove",
		Password: "dasha",
		Dbname:   "film_users",
		Sslmode:  "disable",
	}); err != nil {
		lg.Fatal().Err(err).Msg("faild to initialize database")
	}

	user := &models.User{
		Name:         "undead",
		PasswordHash: "some very strong passord hash",
		Role:         "not admin",
	}

	if id, err := models.NewUser(user.Name, user.PasswordHash, user.Role); err != nil {
		lg.Err(err).Msg("faild to create new user")
	} else {
		lg.Info().Uint("id", id).Msg("new user created")
		user.ID = id
	}

	if isValid, userinfo, err := models.CheckUser(user.Name, user.PasswordHash); err != nil {
		lg.Err(err).Msg("faild to check user")
	} else {
		lg.Info().Bool("is valid", isValid).Str("user info", userinfo.String()).Msg("check correct user")
	}

	if isValid, userinfo, err := models.CheckUser(user.Name, user.PasswordHash+"some wrong hash"); err != nil {
		lg.Err(err).Msg("faild to check user with wrong hash")
	} else {
		lg.Info().Bool("is valid", isValid).Str("user info", userinfo.String()).Msg("check wrong user")
	}

	if userinfo, err := models.GetUserById(user.ID); err != nil {
		lg.Err(err).Msg("faild to get user by id")
	} else {
		lg.Info().Str("user", userinfo.String()).Msg("get user by id")
	}

	if userinfo, err := models.DeleteUserById(user.ID); err != nil {
		lg.Err(err).Msg("faild to delete usser by id")
	} else {
		lg.Info().Str("user", userinfo.String()).Msg("delete user by id")
	}

	lg.Info().Msg("end of the program")
}
