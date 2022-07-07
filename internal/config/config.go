package config

import (
	"sync"

	"github.com/DmitriyZhevnov/rest-api/pkg/logging"
	"github.com/spf13/viper"
)

type Config struct {
	IsDebug *bool `mapstructure:"is_debug"`
	Listen  Listen
	Storage Storage
	Auth    AuthConfig
}

type AuthConfig struct {
	PasswordSalt string
}

type Storage struct {
	Postgresql Postgresql
}

type Postgresql struct {
	Host     string
	Port     string
	Database string
	Username string
	Password string
}

type Listen struct {
	BindIP string `mapstructure:"bind_ip"`
	Port   string `mapstructure:"port"`
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

		if err := setUpViper(); err != nil {
			logger.Fatal(err)
		}

		if err := unmarshal(instance); err != nil {
			logger.Fatal(err)
		}

	})
	return instance
}

func unmarshal(cfg *Config) error {
	if err := viper.Unmarshal(&cfg); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("listen", &cfg.Listen); err != nil {
		return err
	}

	return nil
}

func fromEnv(cfg *Config) error {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	cfg.Auth.PasswordSalt = viper.GetString("PASSWORD_SALT")
	cfg.Storage.Postgresql.Host = viper.GetString("HOST")
	cfg.Storage.Postgresql.Port = viper.GetString("PORT")
	cfg.Storage.Postgresql.Database = viper.GetString("DATABASE")
	cfg.Storage.Postgresql.Username = viper.GetString("USERNAME")
	cfg.Storage.Postgresql.Password = viper.GetString("PASSWORD")

	return nil
}

func setUpViper() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("main")

	return viper.ReadInConfig()
}
