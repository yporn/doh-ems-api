package config

import (
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	AppName     string         `mapstructure:"app_name"`
	Environment string         `mapstructure:"environment"`
	Server      ServerConfig   `mapstructure:"server"`
	Database    DatabaseConfig `mapstructure:"database"`
	MongoDB     MongoDBConfig  `mapstructure:"mongodb"`
}

type ServerConfig struct {
	Host         string `mapstructure:"host"`
	Port         string `mapstructure:"port"`
	ReadTimeout  int    `mapstructure:"read_timeout"`
	WriteTimeout int    `mapstructure:"write_timeout"`
	IdleTimeout  int    `mapstructure:"idle_timeout"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     int `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
	SSLMode  string `mapstructure:"ssl_mode"`
	TimeZone string `mapstructure:"timezone"`
}

type MongoDBConfig struct {
	URI string `mapstructure:"uri"`
}

func (d *DatabaseConfig) GetDSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%d dbname=%s sslmode=%s",
		d.Host, d.Port, d.User, d.Password, d.Name, d.SSLMode)
}

func Load() (*Config, error) {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.AddConfigPath(".")

	viper.SetEnvPrefix("APP")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		var notFound viper.ConfigFileNotFoundError
		if errors.As(err, &notFound) {
			return nil, fmt.Errorf("config file not found: %w", err)
		}
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	if strings.TrimSpace(config.Server.Port) == "" {
		return nil, fmt.Errorf("invalid config: server port is required")
	}

	if config.Environment != "development" {
		if strings.TrimSpace(config.Database.Host) == "" ||
			strings.TrimSpace(config.Database.User) == "" ||
			strings.TrimSpace(config.Database.Name) == "" {
			return nil, fmt.Errorf("invalid config: database configuration is required in %s environment", config.Environment)
		}

		if strings.TrimSpace(config.MongoDB.URI) == "" {
			return nil, fmt.Errorf("invalid config: mongodb URI is required in %s environment", config.Environment)
		}
	}

	return &config, nil
}
