package config

import (
	"os"
)

var BASE_URL string = os.Getenv("BASE_URL")
var BASE_VERSION string = os.Getenv("BASE_VERSION")
var APP_ID string = os.Getenv("APP_ID")
var APP_SECRET string = os.Getenv("APP_SECRET")
var THRESHOLD_BALANCE = 1000 // Threshold balance for the facebook ad account balance for alerting
