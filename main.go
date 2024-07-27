package main

import (
	"log"

	"github.com/MaxSyos/Golang-CRUD/src/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
