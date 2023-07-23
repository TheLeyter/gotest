package main

import (
	"fmt"
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
	Redis       *db.RedisDatabase
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

	if err := app.initRedis(); err != nil {
		return nil, err
	}

	if err := app.initApi(); err != nil {
		return nil, err
	}

	return &app, nil
}

func (app *App) Run() {
	addr := app.Config.API.Host + ":" + app.Config.API.Port
	app.Router.Run(addr)
}

func (app *App) Stop() {
	db, _ := app.DB.DB.DB()
	db.Close()
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

func (app *App) initRedis() error {
	app.Redis = db.NewRedisDatabse(app.Config.Redis)
	app.Redis.Connect()
	err := app.Redis.Redis.Set(ctx, "key", "Hello from go server", 0).Err()
	if err != nil {
		fmt.Println("<<<<<" + err.Error() + ">>>>>>")
	}
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
