package main

import (
	"gotest/common/db"
	"gotest/common/envs"
	"gotest/controllers"
	"gotest/repositories"
	"gotest/services/auth"

	"github.com/gin-gonic/gin"
)

type App struct {
	Router      *gin.Engine
	DB          *db.Database
	Config      *envs.Config
	Controllers []controllers.ApiController
}

func NewApp() (*App, error) {
	app := App{Router: nil, DB: nil, Config: nil}

	if err := app.initConfig(); err != nil {
		return nil, err
	}

	if err := app.initDatabase(); err != nil {
		return nil, err
	}

	if err := app.initApi(); err != nil {
		return nil, err
	}

	return &app, nil
}

func (app *App) Run() {
	app.Router.Run()
}

func (app *App) Stop() {
	// app.Router.St
}

func (app *App) initConfig() error {
	config, err := envs.NewConfig()

	if err != nil {
		return err
	}

	app.Config = config
	return nil
}

func (app *App) initDatabase() error {
	db, err := db.NewDatabase(app.Config.DB)

	if err != nil {
		return err
	}

	if err := db.InitScheme(); err != nil {
		return err
	}

	app.DB = db
	return nil
}

func (app *App) initApi() error {

	engin := gin.Default()
	v1api := engin.Group("/v1")

	userRepository := repositories.NewUserRepository(app.DB)

	authService := auth.NewAuthService(userRepository, app.Config.API)

	app.Controllers = append(app.Controllers, controllers.NewAuthController(v1api, authService))

	for _, s := range app.Controllers {
		s.InitApiRoutes()
	}

	app.Router = engin

	return nil
}
