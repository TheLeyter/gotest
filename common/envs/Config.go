package envs

import (
	"fmt"
	"gotest/common/errors"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type DbConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

type RedisConfig struct {
	Host     string
	Port     string
	User     string
	Password string
}

type ApiConfig struct {
	Secret string
	Host   string
	Port   string
}

type Config struct {
	DB    *DbConfig
	API   *ApiConfig
	Redis *RedisConfig
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

	if err := config.initRedisConfig(); err != nil {
		return nil, err
	}

	fmt.Println(config.Redis)

	return config, nil
}

func (c *Config) initApi() error {
	secret, ok := os.LookupEnv("JWT_SECRET")
	if !ok {
		return &errors.DBError{Message: "JWT_SECRET not found in .env file"}
	}

	host, ok := os.LookupEnv("API_HOST")
	if !ok {
		return &errors.DBError{Message: "API_HOST not found in .env file"}
	}

	port, ok := os.LookupEnv("API_PORT")
	if !ok {
		return &errors.DBError{Message: "API_PORT not found in .env file"}
	}

	c.API = &ApiConfig{Secret: secret, Host: host, Port: port}

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

func (c *Config) initRedisConfig() error {
	host, ok := os.LookupEnv("REDIS_HOST")
	if !ok {
		return &errors.DBError{Message: "REDIS_HOST not found in .env file"}
	}

	user, ok := os.LookupEnv("REDIS_USER")
	if !ok {
		return &errors.DBError{Message: "REDIS_USER not found in .env file"}
	}

	password, ok := os.LookupEnv("REDIS_PASSWORD")
	if !ok {
		return &errors.DBError{Message: "REDIS_PASSWORD not found in .env file"}
	}

	port, ok := os.LookupEnv("REDIS_PORT")
	if !ok {
		return &errors.DBError{Message: "REDIS_PORT not found in .env file"}
	}

	c.Redis = &RedisConfig{Host: host, Port: port, User: user, Password: password}
	return nil
}
