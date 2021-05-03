package main

import (
	"github.com/0xera/app-constructor-backend/api"
	"github.com/0xera/app-constructor-backend/auth"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

func main() {

	if err := initConfig(); err != nil {
		log.Fatal(err)
	}

	jwtService := &auth.JwtService{}
	googleOauthService := auth.CreateService(jwtService)

	apiService := &api.Service{
		JwtService:         jwtService,
		GoogleOauthService: googleOauthService,
	}

	apiService.Serve()
}
func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
