package controllers

import (
	"bytes"
	"encoding/json"
	"log"
	"mh-web-backend/env"
	"mh-web-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func addEmailController(r *gin.Engine, client *http.Client) {
	r.POST("/v1/email/add", func(c *gin.Context) {
		var email models.Email
		c.BindJSON(&email)
		email.Action = "addnoforce"
		postData, err := json.Marshal(email)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			log.Println(err)
			return
		}

		postByteBuffer := bytes.NewBuffer(postData)
		request, err := http.NewRequest("POST", "https://api.mailjet.com/v3/REST/contactslist/1/managecontact", postByteBuffer)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			log.Println(err)
			return
		}

		request.SetBasicAuth(env.MailjetPublicKey, env.MailjetSecretKey)
		request.Header.Set("ContentType", "application/json")

		response, err := client.Do(request)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			log.Println("Failed to send request", err)
			return
		}
		if response.StatusCode < 200 || response.StatusCode >= 300 {
			c.Status(http.StatusUnauthorized)
			log.Println("Mailjet refused request", err)
			defer response.Body.Close()
			return
		}
		defer response.Body.Close()

		c.Status(http.StatusOK)
	})
}

func healthController(r *gin.Engine, client *http.Client) {
	r.GET("/v1/health", func(c *gin.Context) {
		request, err := http.NewRequest("GET", "https://api.mailjet.com/v3/REST/user", nil)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			log.Println(err)
			return
		}
		log.Println(env.MailjetPublicKey, env.MailjetSecretKey)
		request.SetBasicAuth(env.MailjetPublicKey, env.MailjetSecretKey)
		response, err := client.Do(request)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			log.Println("Failed to send request", err)
			return
		}
		if response.StatusCode < 200 || response.StatusCode >= 300 {
			c.Status(http.StatusUnauthorized)
			log.Println("Mailjet refused request", err)
			defer response.Body.Close()
			return
		}
		defer response.Body.Close()

		c.Status(http.StatusOK)
	})
}
