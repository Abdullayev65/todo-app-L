package handler

import (
	"github.com/AbdullohAbdullayev/todo-app-L.git"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary		SignUp
// @Tags			auth
// @Description	create account
// @ID				create-account
// @Accept			json
// @Produce		json
// @Param			input	body		todo.User	true	"account info"
// @Success		200		{integer}	integer		1
// @Failure		400,404	{object}	ResError
// @Failure		500		{object}	ResError
// @Failure		default	{object}	ResError
// @Router			/auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var user todo.User
	if err := c.BindJSON(&user); err != nil {
		newResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.CreateUser(user)
	if err != nil {
		newResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @Summary		SignIn
// @Tags			auth
// @Description	login
// @ID				login
// @Accept			json
// @Produce		json
// @Param			input	body		signInInput	true	"credentials"
// @Success		200		{string}	string		"token"
// @Failure		400,404	{object}	ResError
// @Failure		500		{object}	ResError
// @Failure		default	{object}	ResError
// @Router			/auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var input signInInput
	if err := c.BindJSON(&input); err != nil {
		newResponseError(c, http.StatusBadRequest, err.Error())
		return
	}
	token, err := h.service.GenerateTokenIfExists(input.Username, input.Password)
	if err != nil {
		newResponseError(c, http.StatusConflict, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
