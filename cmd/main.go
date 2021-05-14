package main

import (
	"app-constructor-backend/api"
	"app-constructor-backend/auth"
	"app-constructor-backend/repository"
	"app-constructor-backend/task"
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
	"os"
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
	taskService, err := task.NewTaskServer(repo)
	if err != nil {
		log.Fatal(err)
	}

	socketService := api.NewSocketService(taskService, repo, jwtService)
	if err != nil {
		log.Fatal(err)
	}

	apiService := &api.Service{
		Repository:         repo,
		JwtService:         jwtService,
		GoogleOauthService: googleOauthService,
		SocketService:      socketService,
		TaskService:        taskService,
	}
	if len(os.Args) == 2 && os.Args[1] == "worker" {
		apiService.Serve()
	} else {
		err = taskService.RunWorkers()

	}

	repo.CloseDB()
}
