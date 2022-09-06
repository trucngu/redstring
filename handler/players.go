package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tructn/redstring/common"
	"github.com/tructn/redstring/model"
)

type playerDto struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Ranking   int    `json:"ranking"`
}

func (h *Handler) GetPlayers(c *gin.Context) {
	var players []model.Player
	h.Db.Find(&players)
	fmt.Println(players)
	c.JSON(http.StatusOK, players)
}

func (h *Handler) GetPlayerById(c *gin.Context) {
	var player *model.Player
	var params *common.Params
	if err := c.ShouldBindUri(&params); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	res := h.Db.First(&player, params.Id)
	if res.Error != nil {
		c.JSON(http.StatusBadRequest, res.Error)
		return
	}

	c.JSON(http.StatusOK, player)
}

func (h *Handler) DeletePlayer(c *gin.Context) {
	var params *common.Params
	if err := c.ShouldBindUri(&params); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	h.Db.Delete(&model.Player{}, params.Id)
	c.JSON(http.StatusOK, params.Id)
}

func (h *Handler) CreatePlayer(c *gin.Context) {
	dto := &playerDto{}
	if err := c.BindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	player := &model.Player{
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Email:     dto.Email,
		Ranking:   dto.Ranking,
	}
	res := h.Db.Create(player)

	if res.Error != nil {
		c.JSON(http.StatusBadRequest, "Failed to create player")
		return
	}
	c.JSON(http.StatusOK, dto)
}

func (h *Handler) UpdatePlayer(c *gin.Context) {

	var params *common.Params
	if err := c.ShouldBindUri(&params); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	fmt.Printf("ID = %s", params.Id)

	var dto *playerDto
	if err := c.BindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	repo := &common.Repository[model.Player, string]{Db: h.Db}
	updated, err := repo.Update(params.Id, func(t *model.Player) {
		fmt.Println("Callback function...")
		fmt.Print(t)
		t.FirstName = dto.FirstName
		t.LastName = dto.LastName
		t.Email = dto.Email
		t.Ranking = dto.Ranking
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, updated)
}
