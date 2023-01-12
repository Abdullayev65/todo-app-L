package handler

import (
	"github.com/AbdullohAbdullayev/todo-app-L.git"
	"github.com/AbdullohAbdullayev/todo-app-L.git/pkg/utill"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) creatItem(c *gin.Context) {
	listId, err := utill.ParamInt(c, "id")
	if err != nil {
		newResponseError(c, http.StatusBadRequest, err.Error())
		return
	}
	userId, err := h.getUserId(c)
	if err != nil {
		return
	}

	var intup todo.TodoItem
	if err := c.BindJSON(&intup); err != nil {
		newResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	itemId, err := h.service.TodoItem.Create(userId, listId, intup)
	if err != nil {
		newResponseError(c, http.StatusBadRequest, err.Error())
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, map[string]any{
		"item_id": itemId,
	})
}

func (h *Handler) getAllItem(c *gin.Context) {

}

func (h *Handler) getItemById(c *gin.Context) {

}

func (h *Handler) updateItem(c *gin.Context) {

}

func (h *Handler) deleteItem(c *gin.Context) {

}
