package main

import (
	"mh-web-backend/controllers"
	"mh-web-backend/env"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	env.InitEnvironmentVariables()
	r := gin.Default()
	r.Use(cors.Default())
	client := &http.Client{}

	controllers.InitControllers(r, client)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	r.Run(":" + port)
}
