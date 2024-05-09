package main

import (
	"github.com/gin-gonic/gin"
	"gosweeper/internal/adapters/primary"
	"gosweeper/internal/adapters/secondary"
	"gosweeper/internal/core/services/game_service"
	"gosweeper/pkg"
)

func main() {
	gameRepository := secondary.NewMemKVS()
	gameService := game_service.New(gameRepository, pkg.New())
	gameHandler := primary.NewHttpHandler(gameService)

	router := gin.New()
	router.GET("/games/:id", gameHandler.Get)
	router.POST("/games", gameHandler.Create)

	err := router.Run(":8080")
	if err != nil {
		return
	}
}
