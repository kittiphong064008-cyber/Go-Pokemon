package configs

import (
	"errors"
	"log"

	"github.com/spf13/viper"
)

type Configs struct {
	App        Fiber      `mapstructure:"app"`
	PostgreSQL PostgreSQL `mapstructure:"postgres"`
}

type Fiber struct {
	Host string `mapstructure:"FIBER_HOST"`
	Port string `mapstructure:"FIBER_PORT"`
}

type PostgreSQL struct {
	Host     string `mapstructure:"DB_HOST"`
	Port     string `mapstructure:"DB_PORT"`
	User     string `mapstructure:"DB_USERNAME"`
	Password string `mapstructure:"DB_PASSWORD"`
	Database string `mapstructure:"DB_DATABASE"`
	SSLMode  string `mapstructure:"DB_SSLMODE"`
}

func Load() Configs {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	v.AddConfigPath("./config")

	setDefaults(v)

	if err := v.ReadInConfig(); err != nil {
		var notFound viper.ConfigFileNotFoundError
		if errors.As(err, &notFound) {
			log.Println("Not Found config.yml")
		} else {
			log.Fatalf("can not read config.yml: %v", err)
		}
	}

	var cfg Configs
	if err := v.Unmarshal(&cfg); err != nil {
		log.Fatalf("can not decode config: %v", err)
	}
	return cfg
}

func setDefaults(v *viper.Viper) {
	v.SetDefault("app.FIBER_HOST", "localhost")
	v.SetDefault("app.FIBER_PORT", "8000")
	v.SetDefault("postgres.DB_HOST", "localhost")
	v.SetDefault("postgres.DB_PORT", "5432")
	v.SetDefault("postgres.DB_SSLMODE", "disable")
}
