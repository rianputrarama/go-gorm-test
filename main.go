package main

import (
	"gosqllearn/datasource/mysql"
	"gosqllearn/route"
	"gosqllearn/utils/logging"
)

func main() {
	logging.Init()
	mysql.Init()
	route.Init()
}
