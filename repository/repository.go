package repository

import (
	"app-constructor-backend/model"
	"encoding/base64"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

type Repository struct {
	database    *sqlx.DB
	redisClient *redis.Client
}

func CreateRepository() (*Repository, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		os.Getenv("POSTGRES_NAME"), os.Getenv("DB_PORT"), os.Getenv("DB_USERNAME"),
		os.Getenv("DB_NAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_SSL_MODE")))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)

		return nil, err
	}

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})
	_, err = client.Ping().Result()
	if err != nil {
		log.Fatal(err)

		return nil, err
	}
	return &Repository{database: db, redisClient: client}, nil
}

func (r *Repository) CreateProject(context echo.Context) error {
	project := echo.Map{}
	if err := context.Bind(&project); err != nil {
		return context.String(http.StatusBadRequest, "bind")
	}
	tx, err := r.database.Begin()
	if err != nil {
		return context.String(http.StatusBadRequest, "Begin")
	}
	claims := getUserClaims(context)
	var id int
	row := tx.QueryRow("insert into project(name) values ($1) returning id", project["name"])
	if err := row.Scan(&id); err != nil {
		_ = tx.Rollback()
		return err
	}

	_, err = tx.Exec("insert into user_projects(user_id, project_id) values ($1, $2)", claims.Sub, id)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return context.JSON(http.StatusOK, echo.Map{
		"id": id,
	})
}

func (r *Repository) AddUser(dataJwt *model.UserDataJwt) error {
	_, err := r.database.Exec("insert into user_data(id, email) values ($1, $2) ON CONFLICT DO NOTHING", dataJwt.Sub, dataJwt.Email)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) SaveProject(context echo.Context) error {
	request := &model.RequestBody{}
	if err := context.Bind(request); err != nil {
		return context.String(http.StatusBadRequest, "")
	}

	claims := getUserClaims(context)
	tx, err := r.database.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("update project set app = $1 where id = $2", request.Project.App, request.Project.Id)
	if err != nil {
		return err
	}

	_, err = tx.Exec("update user_data set widgets_count = $2 where id =$2", request.WidgetsCount, claims.Sub)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, echo.Map{
		"ok": true,
	})
}

func (r *Repository) DeleteProject(context echo.Context) error {
	request := echo.Map{}
	if err := context.Bind(&request); err != nil {
		return context.String(http.StatusBadRequest, "")
	}

	_, err := r.database.Exec("delete from project where id = $1", request["id"])
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, echo.Map{
		"ok": true,
	})
}

func (r *Repository) GetProjects(subUser string) ([]model.Project, error) {

	var projects []model.Project

	err := r.database.Select(&projects, "select project.id, project.name, project.app from project inner join user_projects up on project.id = up.project_id where up.user_id = $1", subUser)
	if err != nil {
		return nil, err
	}
	return projects, nil
}

func (r *Repository) GetProject(subUser string, projectId int) (*model.Project, error) {

	project := &model.Project{}

	err := r.database.Get(project, "select project.id, project.name, project.app from project inner join user_projects up on project.id = $1 where up.user_id = $2 limit 1", projectId, subUser)
	if err != nil {
		return nil, err
	}
	return project, nil
}

func (r *Repository) GetWidgetsCount(subUser string) (int, error) {
	var widgetsCount int

	err := r.database.Get(&widgetsCount, "select widgets_count from user_data where id = $1", subUser)
	if err != nil {
		return 0, err
	}
	return widgetsCount, nil
}

func getUserClaims(c echo.Context) *model.UserClaims {
	user := c.Get("user").(*jwt.Token)
	return user.Claims.(*model.UserClaims)
}

func (r *Repository) Restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*model.UserClaims)
	return c.String(http.StatusOK, "Welcome "+claims.Email+"!")
}

func (r *Repository) CloseDB() {
	if err := r.database.Close(); err != nil {
	}
	if err := r.redisClient.Close(); err != nil {
		log.Fatal(err)
	}
}

func (r *Repository) UpdateUserProjects(userSub string, response *model.Response) {
	tx, err := r.database.Begin()
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = tx.Exec("update user_data set widgets_count = $1 where id =$2", response.WidgetsCount, userSub)
	if err != nil {
		fmt.Println(err)
		return
	}
	projects, err := r.GetProjects(userSub)
	if err != nil {
		fmt.Println(err)
		return
	}
	intersection := Intersection(projects, response.Projects)

	for _, projectUpdate := range intersection {
		_, err = tx.Exec("update project set app = $1 where id = $2", projectUpdate.App, projectUpdate.Id)
	}
	except := Except(projects, response.Projects)
	for _, projectDelete := range except {
		_, err = tx.Exec("delete from project where id = $1", projectDelete.Id)

	}
	err = tx.Commit()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (r *Repository) DownloadProject(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*model.UserClaims)
	name := c.Param("name")
	if name == "" {
		return c.JSON(http.StatusBadRequest, "not valid name")
	}

	filePath := filepath.FromSlash("result/" + claims.Sub + "/" + name)

	err := c.File(filePath)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "server cant attach file. please retry")
	}

	return nil
}

func (r *Repository) PublishProject(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*model.UserClaims)
	projectId := c.Param("projectId")
	if projectId == "" {
		return c.JSON(http.StatusBadRequest, "not valid projectId")
	}
	atoi, err := strconv.Atoi(projectId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "not valid projectId")
	}
	project, err := r.GetProject(claims.Sub, atoi)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "not valid projectId")
	}
	baseVal := base64.StdEncoding.EncodeToString([]byte(project.App))
	err = r.redisClient.Set(claims.Sub+projectId, baseVal, 0).Err()

	if err != nil {
		return c.JSON(http.StatusBadRequest, "not valid projectId")
	}
	return c.String(http.StatusOK, "saved")
}

func (r *Repository) ProjectData(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*model.ClientClaims)
	val := r.redisClient.Get(claims.UserId + claims.ProjectId).Val()
	bytes, _ := base64.StdEncoding.DecodeString(val)
	return c.JSONBlob(http.StatusOK, bytes)
}

func Intersection(a, b []model.Project) (c []model.Project) {
	m := make(map[int]bool)

	for _, item := range a {
		m[item.Id] = true
	}

	for _, item := range b {
		if _, ok := m[item.Id]; ok {
			c = append(c, item)
		}
	}
	return c
}
func Except(a, b []model.Project) (c []model.Project) {
	m := make(map[int]bool)

	for _, item := range b {
		m[item.Id] = true
	}

	for _, item := range a {
		if _, ok := m[item.Id]; !ok {
			c = append(c, item)
		}
	}
	return c
}
