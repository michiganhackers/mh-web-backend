package main

import (
	"mh-web-backend/controllers"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	client := &http.Client{}

	controllers.InitControllers(r, client)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	r.Run(":" + port)
}
