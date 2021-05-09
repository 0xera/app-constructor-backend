package task

import (
	"fmt"
	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/RichardKnop/machinery/v1/tasks"
	"os"
	"path/filepath"
	"time"
)

type Service struct {
	server *machinery.Server
}

func (s *Service) RunWorkers() error {
	worker := s.server.NewWorker("build workers", 10)
	if err := worker.Launch(); err != nil {
		return err
	}
	return nil
}

func (s *Service) Execute(userId string, projectId int) (string, error) {
	task := tasks.Signature{
		Name: "build_project",
		Args: []tasks.Arg{
			{
				Type:  "string",
				Value: userId,
			},
			{
				Type:  "int",
				Value: projectId,
			},
		},
	}
	result, err := s.server.SendTask(&task)
	if err != nil {
		return "", err
	}

	value, err := result.Get(time.Second * 5)
	if err != nil {
		return "", err
	}

	newFilename := fmt.Sprintf("%v", value[0].Interface())

	return newFilename, nil
}

func NewTaskServer() (*Service, error) {
	server, err := machinery.NewServer(&config.Config{
		Broker:        "redis://" + os.Getenv("REDIS_PASSWORD") + "@localhost:6379",
		ResultBackend: "redis://" + os.Getenv("REDIS_PASSWORD") + "@localhost:6379",
	})
	if err != nil {
		return nil, err
	}

	err = server.RegisterTasks(map[string]interface{}{
		"build_project": build,
	})
	if err != nil {
		return nil, err
	}

	return &Service{server}, nil
}

func build(userId string, projectId int) (string, error) {
	time.Sleep(time.Second * 20)
	filename := fmt.Sprintf("%d-%d%s", projectId, time.Now().Nanosecond(), ".apk")

	if _, err := os.Stat("result"); os.IsNotExist(err) {
		os.Mkdir("result", os.ModeDir)
	}
	userDir := filepath.FromSlash("result/" + userId)
	if _, err := os.Stat(userDir); os.IsNotExist(err) {
		os.Mkdir(userDir, os.ModeDir)
	}
	_, err := os.Create(filepath.FromSlash(userDir + "/" + filename))
	if err != nil {
		return "", err
	}
	return filename, nil
}
