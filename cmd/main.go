package main

import (
	"app-constructor-backend/api"
	"app-constructor-backend/auth"
	"app-constructor-backend/repository"
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

func main() {

	if err := initConfig(); err != nil {
		log.Fatal(err)
	}

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	jwtService := &auth.JwtService{}
	repo, err := repository.CreateRepository()
	googleOauthService := auth.CreateService(jwtService, repo)

	if err = initConfig(); err != nil {
		log.Fatal(err)
	}

	apiService := &api.Service{
		Repository:         repo,
		JwtService:         jwtService,
		GoogleOauthService: googleOauthService,
	}

	apiService.Serve()

	repo.DestroyDB()
}
func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
