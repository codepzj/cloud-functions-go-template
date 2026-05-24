package conf

import (
	"log"

	"cloud-functions/pkg/logger"

	"github.com/caarlos0/env/v11"
	"github.com/subosito/gotenv"
)

var cfg Config

type Config struct {
	Port     int      `env:"PORT" envDefault:"2574"`
	Log      Log      `envPrefix:"LOG_"`
}

type Log struct {
	Format logger.LogFormat `env:"FORMAT" envDefault:"console"`
	Level  logger.LogLevel  `env:"LEVEL" envDefault:"info"`
	Output LogOutput        `envPrefix:"OUTPUT_"`
}

type LogOutput struct {
	EnableFile bool   `env:"ENABLE_FILE" envDefault:"false"`
	FilePath   string `env:"FILE_PATH" envDefault:"logs/app.log"`
	MaxAge     int    `env:"MAX_AGE" envDefault:"7"`
}

func GetConfig() *Config {
	_ = gotenv.Load()

	var config Config
	if err := env.Parse(&config); err != nil {
		log.Println("parse config from env error:", err.Error())
		panic(err)
	}

	log.Println("config loaded successfully from env")
	cfg = config
	return &cfg
}

func GetGlobalConfig() *Config {
	return &cfg
}
