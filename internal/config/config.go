package config

import (
	"sync"

	"github.com/DmitriyZhevnov/rest-api/pkg/logging"
	"github.com/spf13/viper"
)

type Config struct {
	Storage Storage
	Auth    AuthConfig
}

type AuthConfig struct {
	PasswordSalt string
}

type Storage struct {
	Postgresql Postgresql
	Firestore  Firestore
}

type Firestore struct {
	ProjectID string
}

type Postgresql struct {
	Host                   string
	Port                   string
	Database               string
	Username               string
	Password               string
	DBIAMUser              string
	InstanceConnectionName string
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		logger := logging.GetLogger()
		logger.Info("read application configuration")
		instance = &Config{}

		if err := fromEnv(instance); err != nil {
			logger.Fatal(err)
		}
	})
	return instance
}

func fromEnv(cfg *Config) error {
	if err := viper.BindEnv("PASSWORD_SALT"); err != nil {
		return err
	}

	if err := viper.BindEnv("HOST"); err != nil {
		return err
	}

	if err := viper.BindEnv("DB_PORT"); err != nil {
		return err
	}

	if err := viper.BindEnv("DATABASE"); err != nil {
		return err
	}

	if err := viper.BindEnv("USERNAME"); err != nil {
		return err
	}

	if err := viper.BindEnv("PASSWORD"); err != nil {
		return err
	}

	if err := viper.BindEnv("DB_IAM_USER"); err != nil {
		return err
	}

	if err := viper.BindEnv("INSTANCE_CONNECTION_NAME"); err != nil {
		return err
	}

	if err := viper.BindEnv("PROJECT_ID"); err != nil {
		return err
	}

	cfg.Auth.PasswordSalt = viper.GetString("PASSWORD_SALT")
	cfg.Storage.Postgresql.Host = viper.GetString("HOST")
	cfg.Storage.Postgresql.Port = viper.GetString("DB_PORT")
	cfg.Storage.Postgresql.Database = viper.GetString("DATABASE")
	cfg.Storage.Postgresql.Username = viper.GetString("USERNAME")
	cfg.Storage.Postgresql.Password = viper.GetString("PASSWORD")
	cfg.Storage.Postgresql.DBIAMUser = viper.GetString("DB_IAM_USER")
	cfg.Storage.Postgresql.InstanceConnectionName = viper.GetString("INSTANCE_CONNECTION_NAME")
	cfg.Storage.Firestore.ProjectID = viper.GetString("PROJECT_ID")

	return nil
}
