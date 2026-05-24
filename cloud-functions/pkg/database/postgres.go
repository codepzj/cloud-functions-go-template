package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"server/pkg/logger"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	glogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type PgSQLOptions struct {
	Host             string
	User             string
	Password         string
	DBName           string
	Port             int
	SSLMode          string
	TimeZone         string
	SlowSqlThreshold int
	LogLevel         string
}

var pgSQLGormDB *gorm.DB

func NewPgSQL(opt *PgSQLOptions) *gorm.DB {
	gormLogger := glogger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		glogger.Config{
			SlowThreshold:             time.Duration(opt.SlowSqlThreshold) * time.Millisecond,
			LogLevel:                  parseLogLevel(opt.LogLevel),
			IgnoreRecordNotFoundError: true,
		},
	)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s", opt.Host, opt.User, opt.Password, opt.DBName, opt.Port, opt.SSLMode, opt.TimeZone)
	db, err := gorm.Open(postgres.New(postgres.Config{DSN: dsn}), &gorm.Config{
		Logger: gormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		logger.Error("postgres connect failed", zap.Error(err))
		panic(err)
	}

	logger.Info("postgres connect success...")
	pgSQLGormDB = db
	return db
}

func PgSQLClose() error {
	if pgSQLGormDB != nil {
		sqlDB, err := pgSQLGormDB.DB()
		if err != nil {
			logger.Error("get postgres db failed", zap.Error(err))
			return err
		}
		err = sqlDB.Close()
		if err != nil {
			logger.Error("close postgres db failed", zap.Error(err))
			return err
		}
		logger.Info("close postgres db success...")
		return nil
	}
	return nil
}
