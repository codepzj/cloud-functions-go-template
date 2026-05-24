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
	Postgres Postgres `envPrefix:"POSTGRES_"`
	Log      Log      `envPrefix:"LOG_"`
}

type Postgres struct {
	Host             string `env:"HOST,required"`
	User             string `env:"USER,required"`
	Password         string `env:"PASSWORD,required"`
	DBName           string `env:"DBNAME,required"`
	Port             int    `env:"PORT" envDefault:"5432"`
	SSLMode          string `env:"SSLMODE" envDefault:"disable"`
	TimeZone         string `env:"TIMEZONE" envDefault:"Asia/Shanghai"`
	SlowSqlThreshold int    `env:"SLOW_SQL_THRESHOLD" envDefault:"200"`
	LogLevel         string `env:"LOG_LEVEL" envDefault:"info"`
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
