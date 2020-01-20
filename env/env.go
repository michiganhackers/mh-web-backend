package env

import "os"

var MailjetPublicKey = os.Getenv("MH_MAILJET_PUBLIC")
var MailjetSecretKey = os.Getenv("MH_MAILJET_SECRET")
