package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tructn/redstring/model"
)

type login struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type register struct {
	UserName  string `json:"userName"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (h *Handler) Token(c *gin.Context) {
	var req login
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	user := &model.User{}
	err := h.Db.First(user, "user_name = ?", req.UserName).Error

	if err != nil {
		c.JSON(http.StatusUnauthorized, err)
		return
	}

	passwordMatch := user.VerifyPassword(req.Password)
	if !passwordMatch {
		c.JSON(http.StatusUnauthorized, "Password is invalid")
		return
	}

	fmt.Println("Password matched...")

	jwt, err := user.GenerateToken()
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("Failed to generate JWT %s", err))
		return
	}

	c.JSON(http.StatusOK, jwt)
}

func (h *Handler) Register(c *gin.Context) {
	var req register
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	fmt.Println(req)
	user := &model.User{
		UserName:  req.UserName,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		IsActive:  true,
	}
	user.Hash(req.Password)

	h.Db.Create(&user)

	c.JSON(http.StatusOK, user)
}
