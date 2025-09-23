package main

import (
	"github.com/elyron7/webook/internal/web"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	userHandler := web.NewUserHandler()
	userHandler.RegisterRouter(server)

	err := server.Run() // listens on 0.0.0.0:8080 by default
	if err != nil {
	}
}
