package main

import (
	"app-constructor-backend/api"
	"app-constructor-backend/auth"
	"app-constructor-backend/repository"
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	jwtService := &auth.JwtService{}
	repo, err := repository.CreateRepository()
	googleOauthService := auth.CreateService(jwtService, repo)

	if err != nil {
		log.Fatal(err)
	}

	apiService := &api.Service{
		Repository:         repo,
		JwtService:         jwtService,
		GoogleOauthService: googleOauthService,
	}

	apiService.Serve()

	repo.CloseDB()
}
