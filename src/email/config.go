package email

import "os"

var (
	MJ_APIKEY_PUBLIC  = os.Getenv("MJ_APIKEY_PUBLIC")
	MJ_APIKEY_PRIVATE = os.Getenv("MJ_APIKEY_PRIVATE")
	SENDER_EMAIL      = os.Getenv("SENDER_EMAIL")
	TEMPLATE_ID       = 6267302
)
