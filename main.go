package main

import (
	"exam/config"
	"exam/connection"
	//"exam/connection"
	"exam/repository"
	"exam/server"
	"exam/service"
	"log"

	//"net"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

// @title           Postgres Crud API
// @version         1.0
// @description     This is a sample server celler server.

// NewRoutor
// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io
func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Println("error while configging")
		panic(err)
	}
	db, err := connection.Connect(cfg)
	repo:= repository.New(db)
	svc := service.New(repo)
	svr := server.NewServiceStruct(svc)
	e := gin.Default()
	e.POST("/new-word", svr.TakeWord)
	e.GET("/take-all-word", svr.GiveAllWord)
	e.GET("/take-limited-word", svr.GiveLimitedWord)
	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	e.Run(":9000")
}
