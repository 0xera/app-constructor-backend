package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

type JwtService struct {
	config     middleware.JWTConfig
	claimsType UserClaims
}

type UserDataJwt struct {
	Sub   string `json:"sub"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
type UserClaims struct {
	UserDataJwt
	jwt.StandardClaims
}

func (service JwtService) CreateMiddleware() echo.MiddlewareFunc {
	config := middleware.JWTConfig{
		Claims:     &UserClaims{},
		SigningKey: []byte(viper.GetString("secretJwt")),
	}
	return middleware.JWTWithConfig(config)
}

func (service *JwtService) CreateTokensPair(userDataJwt UserDataJwt) (map[string]string, error) {
	userClaimsAccess := UserClaims{
		UserDataJwt: userDataJwt,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
		},
	}

	userClaimsRefresh := &UserClaims{
		UserDataJwt: userDataJwt,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 70).Unix(),
		},
	}

	var accessToken string
	var refreshToken string
	var err error
	accessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, userClaimsAccess).SignedString([]byte(viper.GetString("secretJwt")))
	if err != nil {
		return nil, err
	}
	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, userClaimsRefresh).SignedString([]byte(viper.GetString("secretJwt")))
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}, nil
}

func (service *JwtService) RefreshToken(c echo.Context) error {
	tokenReq := c.Param("refresh_token")

	if tokenReq == "" {
		return c.JSON(http.StatusBadGateway, "not valid token")
	}

	token, err := jwt.ParseWithClaims(tokenReq, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing")
		}
		return []byte(viper.GetString("secretJwt")), nil
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "token parse error")
	}

	userClaims, ok := token.Claims.(*UserClaims)

	if ok && token.Valid {
		newTokenPair, err := service.CreateTokensPair(UserDataJwt{
			userClaims.Sub,
			userClaims.Name,
			userClaims.Email,
		})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "im done")
		}
		return c.JSON(http.StatusOK, newTokenPair)
	} else {
		return echo.ErrUnauthorized
	}
}
