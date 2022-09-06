package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tructn/redstring/common"
	"github.com/tructn/redstring/env"
	"github.com/tructn/redstring/handler"
)

func main() {
	router := gin.Default()
	env := env.GetEnv()
	handler := &handler.Handler{Db: common.GetDatabase(*env)}

	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", handler.Register)
			auth.POST("/token", handler.Token)
		}

		players := api.Group("/players")
		{
			players.GET("/", handler.GetPlayers)
			players.GET("/:id", handler.GetPlayerById)
			players.POST("/", handler.CreatePlayer)
			players.PUT("/:id", handler.UpdatePlayer)
			players.DELETE("/:id", handler.DeletePlayer)
		}
	}

	router.Run()
}
