package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	client := &http.Client{}

	r.POST("/v1/email/add", func(c *gin.Context) {
		var email EMAIL
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

		var mailjetPublicKey = os.Getenv("MH_MAILJET_PUBLIC")
		var mailjetSecretKey = os.Getenv("MH_MAILJET_SECRET")

		request.SetBasicAuth(mailjetPublicKey, mailjetSecretKey)
		request.Header.Set("ContentType", "application/json")

		response, err := client.Do(request)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			log.Println(err)
			return
		}
		defer response.Body.Close()

		c.Status(http.StatusOK)
	})
	r.Run(":8080")
}
