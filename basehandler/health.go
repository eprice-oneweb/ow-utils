package basehandler

import (
	"net/http"

	"github.com/eprice-oneweb/ow-utils/basemodels"

	"github.com/gin-gonic/gin"
)

// @Summary     Show the status of server.
// @Description get the status of server.
// @Tags        health
// @Accept      */*
// @Produce     json
// @Success     200 {object} models.HealthStatus
// @Router      /health [get]
func Health(c *gin.Context) {
	c.JSON(http.StatusOK, basemodels.HealthStatus{Status: "ok"})
}
