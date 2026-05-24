package handler

import (
	"net/http"
	"server/internal/service"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
	svc *service.HealthService
}

func NewHealthHandler(svc *service.HealthService) *HealthHandler {
	return &HealthHandler{svc: svc}
}

func (h *HealthHandler) Health(c *gin.Context) {
	RespondSuccess(c, http.StatusOK, "ok")
}
