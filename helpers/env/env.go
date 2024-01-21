package env

import (
	"fmt"
	"github.com/spf13/viper"
	"reflect"
)

type Config struct {
	DBHost       string `mapstructure:"DB_HOST" default:"127.0.0.1"`
	DBPort       int    `mapstructure:"DB_PORT" default:"3306"`
	DBName       string `mapstructure:"DB_NAME" default:""`
	DBUser       string `mapstructure:"DB_USER" default:"root"`
	DBPass       string `mapstructure:"DB_PASS" default:""`
	AppHost      string `mapstructure:"APP_HOST" default:"localhost"`
	AppPort      int    `mapstructure:"APP_PORT" default:"5000"`
	DefaultLimit int    `mapstructure:"DEFAULT_LENGTH_PAGINATION" default:"20"`
}

func (c *Config) ConnectionString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.DBUser, c.DBPass, c.DBHost, c.DBPort, c.DBName)
}

var (
	cfg *Config = nil
)

func Get() *Config {
	if cfg == nil {
		cfg = new(Config)

		viper.SetConfigType("yaml")
		viper.SetConfigFile("env.yaml")
		viper.AddConfigPath(".")
		viper.AutomaticEnv()

		_ = viper.ReadInConfig()

		e := reflect.ValueOf(cfg).Elem()
		t := e.Type()
		for i := 0; i < e.NumField(); i++ {
			key := t.Field(i).Tag.Get("mapstructure")
			value := t.Field(i).Tag.Get("default")

			viper.SetDefault(key, value)
		}

		_ = viper.Unmarshal(cfg)
	}

	return cfg
}
