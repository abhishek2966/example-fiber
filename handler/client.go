package handler

import (
	"net/http"
	"time"
)

// var client = http.DefaultClient
var client = &http.Client{
	Timeout: 5 * time.Second,
}
