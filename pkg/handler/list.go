package handler

import (
	"github.com/AbdullohAbdullayev/todo-app-L.git"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) creatList(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		return
	}
	var intup todo.TodoList
	if err := c.BindJSON(&intup); err != nil {
		newResponseError(c, http.StatusBadRequest, err.Error())
		return
	}
	listId, err := h.service.Create(userId, intup)
	if err != nil {
		newResponseError(c, http.StatusBadRequest, err.Error())
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, map[string]any{
		"list_id": listId,
	})
}

type getAllListsResponse struct {
	Data []todo.TodoList
}

func (h *Handler) getAllList(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		return
	}
	lists, err := h.service.GetAll(userId)
	if err != nil {
		newResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.AbortWithStatusJSON(statusOk, getAllListsResponse{Data: lists})
}
func (h *Handler) getListById(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		return
	}
	var id int
	if paramId := c.Param("id"); paramId == "" {
		newResponseError(c, http.StatusBadRequest, "param id not founded")
		return
	} else {
		if i, err := strconv.Atoi(paramId); err != nil {
			newResponseError(c, http.StatusBadRequest, "param id invalid")
			return
		} else {
			id = i
		}
	}
	list, err := h.service.GetById(userId, id)
	if err != nil {
		newResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.AbortWithStatusJSON(statusOk, list)
}
func (h *Handler) updateList(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		return
	}
	var id int
	if paramId := c.Param("id"); paramId == "" {
		newResponseError(c, http.StatusBadRequest, "param id not founded")
		return
	} else {
		if i, err := strconv.Atoi(paramId); err != nil {
			newResponseError(c, http.StatusBadRequest, "param id invalid")
			return
		} else {
			id = i
		}
	}
	var input todo.InputListUpdate
	c.BindJSON(&input)
	err = h.service.Update(userId, id, input)
	if err != nil {
		newResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.AbortWithStatusJSON(statusOk, newStatusRes("ok"))
}
func (h *Handler) deleteList(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		return
	}
	var id int
	if paramId := c.Param("id"); paramId == "" {
		newResponseError(c, http.StatusBadRequest, "param id not founded")
		return
	} else {
		if i, err := strconv.Atoi(paramId); err != nil {
			newResponseError(c, http.StatusBadRequest, "param id invalid")
			return
		} else {
			id = i
		}
	}
	err = h.service.Delete(userId, id)
	if err != nil {
		newResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.AbortWithStatusJSON(statusOk, newStatusRes("ok"))
}
