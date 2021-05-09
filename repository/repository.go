package repository

import (
	"app-constructor-backend/model"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type Repository struct {
	database *sqlx.DB
}

func CreateRepository() (*Repository, error) {
	value := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("POSTGRES_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSL_MODE"))
	db, err := sqlx.Open("postgres", value)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)

		return nil, err
	}
	return &Repository{db}, nil
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
	err = os.Remove(filePath)
	if err != nil {
		fmt.Println(err)
	}
	return nil
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
