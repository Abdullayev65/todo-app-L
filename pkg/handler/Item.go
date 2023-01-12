package handler

import (
	"github.com/AbdullohAbdullayev/todo-app-L.git"
	"github.com/AbdullohAbdullayev/todo-app-L.git/pkg/utill"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) creatItem(c *gin.Context) {
	listId, err := utill.ParamInt(c, "list_id")
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

type getAllItemsResponse struct {
	Data []todo.TodoItem
}

func (h *Handler) getAllItem(c *gin.Context) {
	listId, err := utill.ParamInt(c, "list_id")
	if err != nil {
		newResponseError(c, http.StatusBadRequest, err.Error())
		return
	}
	userId, err := h.getUserId(c)
	if err != nil {
		newResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	lists, err := h.service.TodoItem.GetAll(userId, listId)
	if err != nil {
		newResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.AbortWithStatusJSON(statusOk, getAllItemsResponse{Data: lists})
}

func (h *Handler) getItemById(c *gin.Context) {
	listId, err := utill.ParamInt(c, "list_id")
	if err != nil {
		newResponseError(c, http.StatusBadRequest, err.Error())
		return
	}
	userId, err := h.getUserId(c)
	if err != nil {
		return
	}
	itemId, err := utill.ParamInt(c, "id")
	if err != nil {
		newResponseError(c, http.StatusBadRequest, err.Error())
		return
	}
	list, err := h.service.TodoItem.GetById(userId, listId, itemId)
	if err != nil {
		newResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.AbortWithStatusJSON(statusOk, list)
}

func (h *Handler) updateItem(c *gin.Context) {
	listId, err := utill.ParamInt(c, "list_id")
	if err != nil {
		newResponseError(c, http.StatusBadRequest, err.Error())
		return
	}
	userId, err := h.getUserId(c)
	if err != nil {
		return
	}
	itemId, err := utill.ParamInt(c, "id")
	if err != nil {
		newResponseError(c, http.StatusBadRequest, err.Error())
		return
	}
	var input todo.InputItemUpdate
	c.BindJSON(&input)
	err = h.service.TodoItem.Update(userId, listId, itemId, input)
	if err != nil {
		newResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.AbortWithStatusJSON(statusOk, newStatusRes("ok"))
}

func (h *Handler) deleteItem(c *gin.Context) {
	listId, err := utill.ParamInt(c, "list_id")
	if err != nil {
		newResponseError(c, http.StatusBadRequest, err.Error())
		return
	}
	userId, err := h.getUserId(c)
	if err != nil {
		return
	}
	itemId, err := utill.ParamInt(c, "id")
	if err != nil {
		newResponseError(c, http.StatusBadRequest, err.Error())
		return
	}
	err = h.service.TodoItem.Delete(userId, listId, itemId)
	if err != nil {
		newResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.AbortWithStatusJSON(statusOk, newStatusRes("ok"))
}
