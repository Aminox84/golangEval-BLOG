package main

import (
	"blog/internal/database"
	"blog/web"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("ok")

	database.Init()

	router := gin.Default()

	server := web.NewServer(router)
	server.RegisterRoutes()

	if err := server.Run(":8080"); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
