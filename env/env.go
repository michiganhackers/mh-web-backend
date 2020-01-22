package env

import (
	"log"
	"os"
)

var MailjetPublicKey, MailjetSecretKey string

func InitEnvironmentVariables() {
	var ok bool
	MailjetPublicKey, ok = os.LookupEnv("MH_MAILJET_PUBLIC")
	if !ok {
		log.Println("MH_MAILJET_PUBLIC env not present")
	}
	MailjetSecretKey, ok = os.LookupEnv("MH_MAILJET_SECRET")
	if !ok {
		log.Println("MH_MAILJET_SECRET env not present")
	}
}
