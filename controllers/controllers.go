package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// InitControllers attaches the controllers (listeners) for HTTP requests to the Gin engine.
func InitControllers(r *gin.Engine, client *http.Client) {
	addEmailController(r, client)
	healthController(r, client)
}
