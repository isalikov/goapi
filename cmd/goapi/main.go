// Package goapi
//
// Goapi endpoints
//
//     Schemes: http
//     BasePath: /
//     Version: 1.0.0
//
//     Produces:
//     - application/json
//
// swagger:meta
package main

import (
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/isalikov/goapi/internal/env"
	"github.com/isalikov/goapi/internal/services"
	"github.com/isalikov/goapi/internal/utils"
	"github.com/isalikov/goapi/routes/home"
	"log"
)


func main() {
	envConfig := &env.Config{}

	err := env.Parse(envConfig)
	if err != nil {
		log.Fatalln(err, "Parsing environment")
	}

	logger := &utils.Logger{Debug: envConfig.Env != "release"}
	logger.Init(envConfig.SentryDsn, envConfig.Release)

	clients := services.Setup(envConfig, logger)

	r := gin.Default()
	ctx := context.Background()

	r.Static("/docs", "swagger-ui")

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
	}))

	r.GET("/home/example", home.Example(ctx, clients, logger))

	err = r.Run(fmt.Sprintf("0.0.0.0:%v", envConfig.Port))
	if err != nil {
		logger.Fatal(err, "Gin instance error")
	}
}
