package configs

import (
    "os"
)

var MYSQL_HOST = os.Getenv("MYSQL_HOST")
var MYSQL_PASSWORD = os.Getenv("MYSQL_PASSWORD")
var MYSQL_PORT = os.Getenv("MYSQL_PORT")
var MYSQL_USERNAME = os.Getenv("MYSQL_USER")
const MYSQL_NAME = "tiktok_backend"
