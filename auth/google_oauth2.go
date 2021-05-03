package auth

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io"
	"net/http"
)

type GoogleOauthService struct {
	JwtService  *JwtService
	oauthConfig *oauth2.Config
}

func CreateService(service *JwtService) *GoogleOauthService {
	return &GoogleOauthService{
		service,
		&oauth2.Config{
			ClientID:     viper.GetString("clientId"),
			ClientSecret: viper.GetString("clientSecret"),
			RedirectURL:  viper.GetString("redirectUrl"),
			Scopes: []string{
				"https://www.googleapis.com/auth/userinfo.email",
				"https://www.googleapis.com/auth/userinfo.profile",
			},
			Endpoint: google.Endpoint,
		},
	}
}

func stateGenerator() string {
	b := make([]byte, 4)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func (service *GoogleOauthService) AuthUrl(echo echo.Context) error {
	return echo.Redirect(http.StatusPermanentRedirect, service.oauthConfig.AuthCodeURL(stateGenerator()))
}

func (service *GoogleOauthService) Login(echoContext echo.Context) error {
	codeParam := echoContext.QueryParam("code")
	token, err := service.oauthConfig.Exchange(context.Background(), codeParam)
	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, err)
	}

	userDataJwt, err := service.requestGoogleProfileData(token)
	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, err)
	}

	newTokenPair, err := service.JwtService.CreateTokensPair(*userDataJwt)
	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, err)
	}

	return echoContext.JSON(http.StatusOK, newTokenPair)
}

func (service *GoogleOauthService) requestGoogleProfileData(token *oauth2.Token) (*UserDataJwt, error) {
	client := service.oauthConfig.Client(context.Background(), token)
	res, err := client.Get("https://www.googleapis.com/oauth2/v1/userinfo?alt=json")
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(res.Body)

	userDataJwt := &UserDataJwt{}
	err = json.NewDecoder(res.Body).Decode(&userDataJwt)
	if err != nil {
		return nil, err
	}
	return userDataJwt, nil
}
