package config

import (
    "os"
)

/*
const MYSQL_HOST = "172.16.32.82"
const MYSQL_PASSWORD = "PtoYbtCt"
const MYSQL_PORT = "49408"
const MYSQL_USERNAME = "root"
const MYSQL_NAME = "tiktok_backend"
*/

var MYSQL_HOST = os.Getenv("MYSQL_HOST")
var MYSQL_PASSWORD = os.Getenv("MYSQL_PASSWORD")
var MYSQL_PORT = os.Getenv("MYSQL_PORT")
var MYSQL_USERNAME = os.Getenv("MYSQL_USER")
const MYSQL_NAME = "tiktok_backend"

