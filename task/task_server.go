package task

import (
	"app-constructor-backend/model"
	"app-constructor-backend/repository"
	"app-constructor-backend/task/pb"
	"fmt"
	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/RichardKnop/machinery/v1/tasks"
	"google.golang.org/protobuf/encoding/protojson"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

type Service struct {
	server     *machinery.Server
	repository *repository.Repository
}

func (s *Service) RunWorkers() error {
	worker := s.server.NewWorker("build workers", 10)
	if err := worker.Launch(); err != nil {
		return err
	}
	return nil
}

func (s *Service) Execute(userId string, projectId int) (string, error) {
	project, err := s.repository.GetProject(userId, projectId)

	if err != nil {
		return "", err
	}
	bytes, err := project.MarshalJSON()

	if err != nil {
		return "", err
	}

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
			{
				Type:  "string",
				Value: string(bytes),
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

func NewTaskServer(repository *repository.Repository) (*Service, error) {
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

	return &Service{server, repository}, nil
}

func build(userId string, projectId int, project string) (string, error) {

	fakeId := fmt.Sprintf("%d", time.Now().Nanosecond())
	templatesDir := filepath.FromSlash("templates/" + userId + "/" + fakeId)
	if _, err := os.Stat(templatesDir); os.IsNotExist(err) {
		err := os.MkdirAll(templatesDir, os.ModeDir)
		if err != nil {
			return "", err
		}
	}

	p := &model.Project{}
	err := p.UnmarshalJSON([]byte(project))

	err = copyDirectory(templatesDir)
	if err != nil {
		err = os.RemoveAll(templatesDir)
		return "", err
	}
	iconGenerator := &IconGenerator{}

	app := &pb.App{}

	err = protojson.Unmarshal([]byte(p.App), app)
	if err != nil {
		err = os.RemoveAll(templatesDir)
		return "", err
	}
	err = iconGenerator.generateIcon(app, templatesDir)
	if err != nil {
		err = os.RemoveAll(templatesDir)
		return "", err
	}

	stringsResPath := filepath.FromSlash(templatesDir + "/AppConstructor/app/src/main/res/values/strings.xml")
	read, err := ioutil.ReadFile(stringsResPath)
	if err != nil {
		return "", err
	}

	newContents := strings.Replace(string(read), "AppConstructor", p.Name, -1)

	fmt.Println(newContents)

	err = ioutil.WriteFile(stringsResPath, []byte(newContents), 0)
	if err != nil {
		return "", err
	}

	stringsResPath = filepath.FromSlash(templatesDir + "/AppConstructor/app/src/build.gradle")
	read, err = ioutil.ReadFile(stringsResPath)
	if err != nil {
		return "", err
	}

	newContents = strings.Replace(string(read), "com.app.constructor", "com.app.constructor"+fakeId+userId, -1)

	fmt.Println(newContents)

	err = ioutil.WriteFile(stringsResPath, []byte(newContents), 0)
	if err != nil {
		return "", err
	}

	userDir := filepath.FromSlash("result/" + userId)
	if _, err := os.Stat(userDir); os.IsNotExist(err) {
		err := os.Mkdir(userDir, os.ModeDir)
		if err != nil {
			err = os.RemoveAll(templatesDir)
			return "", err
		}
	}
	filename := p.Name + fakeId + ".apk"

	projectDir := filepath.FromSlash(templatesDir + "/AppConstructor")

	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", filepath.FromSlash("task/task"), projectDir)
	} else {
		cmd = exec.Command(filepath.FromSlash("./task/task.sh"), projectDir)
	}
	err = cmd.Run()
	if err != nil {
		err = os.RemoveAll(templatesDir)
		return "", err
	}

	err = CopyFile(filepath.FromSlash(projectDir+"/app/build/outputs/apk/debug/app-debug.apk"), filepath.FromSlash(userDir+"/"+filename))
	if err != nil {
		return "", err
	}
	err = os.RemoveAll(templatesDir)
	if err != nil {
		fmt.Println(err)
	}
	return filename, nil
}

func copyDirectory(dir string) error {
	err := CopyDir("AppConstructor", filepath.FromSlash(dir+"/AppConstructor"))
	if err != nil {
		return err
	}
	return nil
}
