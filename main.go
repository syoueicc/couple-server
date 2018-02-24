package main

import (
	"github.com/cancelv5/couple-server/api"
)

func main() {

	router := router.InitRouter()
	router.Run(":8080")
}
