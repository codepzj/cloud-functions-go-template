package main

import (
	"fmt"
	"os"
	"os/signal"
	"server/conf"
	"server/internal/handler"
	"server/internal/service"
	"server/pkg/database"
	"server/pkg/logger"
	"syscall"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	// 加载配置
	config := conf.GetConfig()

	// 初始化基础设施
	logger.NewLogger(&logger.Option{
		Format: config.Log.Format,
		Level:  config.Log.Level,
		Output: &logger.OutputConfig{
			EnableFile: config.Log.Output.EnableFile,
			FilePath:   config.Log.Output.FilePath,
			MaxAge:     config.Log.Output.MaxAge,
		},
	})
	defer logger.Sync()
	// 初始化数据库
	_ = database.NewPgSQL(&database.PgSQLOptions{
		Host:             config.Postgres.Host,
		User:             config.Postgres.User,
		Password:         config.Postgres.Password,
		DBName:           config.Postgres.DBName,
		Port:             config.Postgres.Port,
		SSLMode:          config.Postgres.SSLMode,
		TimeZone:         config.Postgres.TimeZone,
		SlowSqlThreshold: config.Postgres.SlowSqlThreshold,
		LogLevel:         config.Postgres.LogLevel,
	})

	healthSvc := service.NewHealthService()
	healthHandler := handler.NewHealthHandler(healthSvc)

	// 注册路由
	r := gin.Default()
	r.GET("/health", healthHandler.Health)

	go func() {
		r.Run(fmt.Sprintf(":%d", config.Port))
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	logger.Info("shutting down server...")

	// 关闭数据库连接
	if err := database.PgSQLClose(); err != nil {
		logger.Error("failed to close database", zap.Error(err))
	}
	logger.Info("server shutdown gracefully...")

	os.Exit(0)
}
