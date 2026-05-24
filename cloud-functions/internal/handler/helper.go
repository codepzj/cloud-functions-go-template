package handler

import (
	"errors"
	"net/http"

	"cloud-functions/internal/model"
	"cloud-functions/pkg/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// RespondSuccess 统一的成功响应格式
func RespondSuccess(c *gin.Context, code int, data any) {
	c.JSON(http.StatusOK, model.ApiResponse{
		Success: true,
		Code:    code,
		Error:   "",
		Data:    data,
	})
}

// RespondError 统一的错误响应格式
func RespondError(c *gin.Context, code int, errMsg string) {
	c.JSON(http.StatusOK, model.ApiResponse{
		Success: false,
		Code:    code,
		Error:   errMsg,
	})
}

func RespondBusinessError(c *gin.Context, err error, data any) {
	var bizErr *model.BusinessError
	if errors.As(err, &bizErr) {
		// 业务错误(如密码错误等), 返回具体错误信息
		c.JSON(http.StatusOK, model.ApiResponse{
			Success: false,
			Code:    bizErr.Code,
			Error:   bizErr.Message,
			Data:    data,
		})
	} else {
		// 系统错误, 记录日志并返回通用错误
		logger.Error("Business error", zap.Error(err), zap.String("path", c.Request.URL.Path))
		c.JSON(http.StatusInternalServerError, model.ApiResponse{
			Success: false,
			Code:    http.StatusInternalServerError,
			Error:   "Internal server error",
		})
	}
}
