package tests

import (
	"net/http"
	"os"
)

var (
	publicKey    = os.Getenv("PUBLIC_KEY")
	secretKey    = os.Getenv("SECRET_KEY")
	clientHeader = http.Header{
		"Content-Type":  []string{"application/json"},
		"Authorization": []string{secretKey},
		"x-api-key":     []string{publicKey},
	}
)
