package auth

import (
	"app-constructor-backend/model"
	"app-constructor-backend/repository"
	"context"
	"crypto/rand"
	"fmt"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type GoogleOauthService struct {
	JwtService  *JwtService
	repo        *repository.Repository
	oauthConfig *oauth2.Config
}

func CreateService(service *JwtService, repo *repository.Repository) *GoogleOauthService {
	return &GoogleOauthService{
		service,
		repo,
		&oauth2.Config{
			ClientID:     os.Getenv("CLIENT_ID"),
			ClientSecret: os.Getenv("CLIENT_SECRET"),
			RedirectURL:  os.Getenv("REDIRECT_URL"),
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
	return echo.JSON(http.StatusOK,
		map[string]string{
			"authUrl": service.oauthConfig.AuthCodeURL(stateGenerator()),
		})
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

	err = service.repo.AddUser(userDataJwt)

	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, err)
	}

	return echoContext.JSON(http.StatusOK, newTokenPair)
}

func (service *GoogleOauthService) requestGoogleProfileData(token *oauth2.Token) (*model.UserDataJwt, error) {
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

	userDataJwt := &model.UserDataJwt{}
	bytes, err := ioutil.ReadAll(res.Body)

	err = userDataJwt.UnmarshalJSON(bytes)
	if err != nil {
		return nil, err
	}
	return userDataJwt, nil
}
