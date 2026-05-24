package main

import (
	"cloud-functions/conf"
	"cloud-functions/internal/handler"
	"cloud-functions/internal/service"
	"cloud-functions/pkg/database"
	"cloud-functions/pkg/logger"
	"fmt"
	"os"
	"os/signal"
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

	healthSvc := service.NewHealthService()
	healthHandler := handler.NewHealthHandler(healthSvc)

	// 注册路由
	r := gin.Default()
	r.GET("/api/health", healthHandler.Health)

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
