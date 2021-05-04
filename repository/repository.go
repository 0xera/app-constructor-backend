package repository

import (
	"app-constructor-backend/model"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
)

type Repository struct {
	database *sqlx.DB
}

func CreateRepository() (*Repository, error) {
	value := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		viper.GetString("db.username"),
		os.Getenv("DB_PASSWORD"),
		viper.GetString("db.host"),
		viper.GetString("db.port"),
		viper.GetString("db.name"),
		viper.GetString("db.ssl"))
	db, err := sqlx.Open("postgres", value)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &Repository{db}, nil
}

func (r *Repository) CreateProject(context echo.Context) error {
	project := &model.Project{}
	if err := context.Bind(project); err != nil {
		return context.String(http.StatusBadRequest, "bind")
	}
	tx, err := r.database.Begin()
	if err != nil {
		return context.String(http.StatusBadRequest, "Begin")
	}
	claims := getUserClaims(context)
	var id int
	row := tx.QueryRow("insert into project(id, name) values ($1, $2) returning _id", project.Id, project.Name)
	if err := row.Scan(&id); err != nil {
		_ = tx.Rollback()
		return context.String(http.StatusBadRequest, "create1")
	}

	_, err = tx.Exec("insert into user_projects(user_id, project_id) values ($1, $2)", claims.Sub, id)
	if err != nil {
		return context.String(http.StatusBadRequest, "create2")
	}
	err = tx.Commit()
	if err != nil {
		return context.String(http.StatusBadRequest, "commit")
	}
	return context.JSON(http.StatusOK, echo.Map{
		"isCreate": true,
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
	_, err = tx.Exec("update project set app = $1 where _id = $2", request.Project.App, claims.Sub)
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
	if err := context.Bind(request); err != nil {
		return context.String(http.StatusBadRequest, "")
	}

	claims := getUserClaims(context)
	tx, err := r.database.Begin()
	if err != nil {
		return err
	}
	var id int
	row := tx.QueryRow("select _id from project inner join user_projects up on project._id = up.project_id where up.user_id = $1 and id = $2;", claims.Sub, request["id"])
	if err := row.Scan(&id); err != nil {
		err := tx.Rollback()
		return err
	}

	_, err = tx.Exec("delete from project where _id = $1", id)
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

func (r *Repository) GetProjects(context echo.Context) error {

	var projects []model.Project

	claims := getUserClaims(context)

	err := r.database.Select(&projects, "select project.id, project.name, project.app from project inner join user_projects up on project._id = up.project_id where up.user_id = $1", claims.Sub)
	if err != nil {
	}
	var widgetsCount int

	err = r.database.Get(&widgetsCount, "select widgets_count from user_data where id = $1", claims.Sub)
	if err != nil {
	}

	return context.JSON(http.StatusOK, model.Response{
		WidgetsCount: widgetsCount,
		Projects:     projects,
	})
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

func (r *Repository) DestroyDB() {
	if err := r.database.Close(); err != nil {
		log.Fatal(err)
	}
}
