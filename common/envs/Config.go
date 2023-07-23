package envs

import (
	"gotest/common/errors"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type DbConfig struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     int
}

type ApiConfig struct {
	Secret string
}

type Config struct {
	DB  *DbConfig
	API *ApiConfig
}

func NewConfig() (*Config, error) {
	config := &Config{}

	if err := godotenv.Load(".env"); err != nil {
		return nil, err
	}

	if err := config.initDbConfig(); err != nil {
		return nil, err
	}

	if err := config.initApi(); err != nil {
		return nil, err
	}

	return config, nil
}

func (c *Config) initApi() error {
	secret, ok := os.LookupEnv("JWT_SECRET")
	if !ok {
		return &errors.DBError{Message: "JWT_SECRET not found in .env file"}
	}

	c.API = &ApiConfig{Secret: secret}

	return nil
}

func (c *Config) initDbConfig() error {
	dbHost, ok := os.LookupEnv("POSTGRES_HOST")
	if !ok {
		return &errors.DBError{Message: "POSTGRES_HOST not found in .env file"}
	}

	dbUser, ok := os.LookupEnv("POSTGRES_USER")
	if !ok {
		return &errors.DBError{Message: "POSTGRES_USER not found in .env file"}
	}

	dbPassword, ok := os.LookupEnv("POSTGRES_PASSWORD")
	if !ok {
		return &errors.DBError{Message: "POSTGRES_PASSWORD not found in .env file"}
	}

	dbName, ok := os.LookupEnv("POSTGRES_DB")
	if !ok {
		return &errors.DBError{Message: "POSTGRES_DB not found in .env file"}
	}

	dbPortString, ok := os.LookupEnv("POSTGRES_PORT")
	if !ok {
		return &errors.DBError{Message: "POSTGRES_PORT not found in .env file"}
	}

	dbPort, error := strconv.Atoi(dbPortString)

	if error != nil {
		return error
	}

	c.DB = &DbConfig{Host: dbHost, User: dbUser, Password: dbPassword, Name: dbName, Port: dbPort}

	return nil
}
