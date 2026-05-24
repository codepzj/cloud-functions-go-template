package handler

import (
	"cloud-functions/internal/service"
	"net/http"

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
