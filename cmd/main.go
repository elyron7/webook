package main

import (
	"log"
	"strings"
	"time"

	"github.com/elyron7/webook/internal/repository"
	"github.com/elyron7/webook/internal/repository/dao"
	"github.com/elyron7/webook/internal/service"
	"github.com/elyron7/webook/internal/web"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	server := initServer()
	db := initDB()
	u := initUser(db)
	u.RegisterRouter(server)

	if err := server.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func initDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/webook"))
	if err != nil {
		panic("failed to connect database")
	}
	err = dao.InitTable(db)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func initUser(db *gorm.DB) *web.UserHandler {
	userDAO := dao.NewUserDAO(db)
	userRepo := repository.NewUserRepository(userDAO)
	userService := service.NewUserService(userRepo)
	userHandler := web.NewUserHandler(userService)

	return userHandler
}

func initServer() *gin.Engine {
	server := gin.Default()

	server.Use(cors.New(cors.Config{
		//AllowOrigins:     []string{"http://localhost:3000"},
		//AllowMethods:     []string{"PUT", "PATCH", "GET", "POST"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://localhost") {
				return true
			}
			return strings.Contains(origin, "company.com")
		},
		MaxAge: 12 * time.Hour,
	}))
	return server
}
