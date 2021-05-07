package auth

import (
	"app-constructor-backend/model"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"os"
	"time"
)

var (
	tokens = make(map[string]bool)
)

type JwtService struct {
	config     middleware.JWTConfig
	claimsType model.UserClaims
}

func (service JwtService) CreateMiddleware() echo.MiddlewareFunc {
	config := middleware.JWTConfig{
		Claims:     &model.UserClaims{},
		SigningKey: []byte(os.Getenv("SECRET_JWT")),
	}
	return middleware.JWTWithConfig(config)
}

func (service *JwtService) CreateTokensPair(userDataJwt model.UserDataJwt) (map[string]string, error) {
	var accessToken string
	var refreshToken string
	var err error

	accessToken, err = service.createToken(userDataJwt, os.Getenv("SECRET_JWT"), time.Minute*30)
	if err != nil {
		return nil, err
	}

	refreshToken, err = service.createToken(userDataJwt, os.Getenv("SECRET_JWT_REFRESH"), time.Hour*70)
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	}, nil
}

func (service *JwtService) createToken(userDataJwt model.UserDataJwt, secretKey string, duration time.Duration) (string, error) {
	userClaimsRefresh := &model.UserClaims{
		UserDataJwt: userDataJwt,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(duration).Unix(),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaimsRefresh).SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (service *JwtService) RefreshToken(c echo.Context) error {
	tokenReq := c.Param("refreshToken")

	if tokenReq == "" {
		return c.JSON(http.StatusBadGateway, "not valid token")
	}

	token, err := service.ParseToken(tokenReq, os.Getenv("SECRET_JWT_REFRESH"))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "token parse error")
	}

	userClaims, ok := token.Claims.(*model.UserClaims)

	if ok && token.Valid {
		newTokenPair, err := service.CreateTokensPair(model.UserDataJwt{
			Sub:   userClaims.Sub,
			Name:  userClaims.Name,
			Email: userClaims.Email,
		})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "im done")
		}
		return c.JSON(http.StatusOK, newTokenPair)
	} else {
		return echo.ErrUnauthorized
	}
}

func (service *JwtService) ParseToken(tokenReq string, secretKey string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenReq, &model.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing")
		}
		return []byte(secretKey), nil
	})
	return token, err
}

func (service *JwtService) SocketToken(context echo.Context) error {
	user := context.Get("user").(*jwt.Token)
	userClaims := user.Claims.(*model.UserClaims)
	userDataJwt := model.UserDataJwt{
		Sub:   userClaims.Sub,
		Name:  userClaims.Name,
		Email: userClaims.Email,
	}
	token, err := service.createToken(userDataJwt, os.Getenv("SECRET_JWT_SOCKET"), time.Minute)
	if err != nil {
		return err
	}

	tokens[token] = true
	fmt.Println(token)
	return context.JSON(http.StatusOK, echo.Map{
		"token": token,
	})

}

func (service *JwtService) ContainsToken(token string) bool {
	_, ok := tokens[token]
	return ok
}

func (service *JwtService) DeleteToken(token string) {
	delete(tokens, token)
}
