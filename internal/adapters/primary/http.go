package primary

import (
	"github.com/gin-gonic/gin"
	"gosweeper/internal/core/ports"
	"net/http"
)

type HttpHandler struct {
	gameService ports.GameService
}

func NewHttpHandler(gameService ports.GameService) *HttpHandler {
	return &HttpHandler{
		gameService: gameService,
	}
}

func (handler *HttpHandler) Get(c *gin.Context) {
	game, err := handler.gameService.Get(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, game)
}

func (handler *HttpHandler) Create(c *gin.Context) {
	body := BodyCreate{}
	err := c.Bind(&body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	game, err := handler.gameService.Create(body.Name, body.Size, body.Bombs)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, BuildResponseCreate(game))
}

func (handler *HttpHandler) Reveal(c *gin.Context) {
	body := BodyRevealCell{}
	err := c.Bind(&body)
	id := c.Param("id")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	game, err := handler.gameService.Reveal(id, body.Row, body.Col)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, BuildResponseCreate(game))
}
