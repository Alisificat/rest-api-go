package handler

import (
	todo "github.com/Serminaz/GoRun/todo2"
	"github.com/gin-gonic/gin"
	"net/http"
)

// пустые обработчики для авторизации или регистрации
func (h *Handler) signUp(c *gin.Context) { //объект
	var input todo.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	//	h.services.Authorization
	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	//	h.services.Authorization
	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
