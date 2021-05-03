package repository

import (
	"app-constructor-backend/auth"
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
	"net/http"
	"os"
)

type Repository struct {
	db *sqlx.DB
}

var ctx = context.Background()

func createRepository() (*Repository, error) {
	db, err := sqlx.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
			viper.GetString("db.host"),
			viper.GetString("db.port"),
			viper.GetString("db.name"),
			viper.GetString("db.name"),
			os.Getenv("DB_PASSWORD"),
			viper.GetString("db.ssl")))

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &Repository{db}, nil
}

func (repository *Repository) SaveProject(context echo.Context) error {
	project := &Project{}
	if err := context.Bind(project); err != nil {
		return context.String(http.StatusBadRequest, "")
	}
	return nil
	//
	//claims := getUserClaims(context)
	//
	//repository.db.Query()
}

func getUserClaims(c echo.Context) *auth.UserClaims {
	user := c.Get("user").(*jwt.Token)
	return user.Claims.(*auth.UserClaims)
}
