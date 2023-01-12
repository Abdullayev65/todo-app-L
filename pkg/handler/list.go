package handler

import (
	"github.com/AbdullohAbdullayev/todo-app-L.git"
	"github.com/AbdullohAbdullayev/todo-app-L.git/pkg/utill"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary		Create todo list
// @Security		ApiKeyAuth
// @Tags			lists
// @Description	create todo list
// @ID				create-list
// @Accept			json
// @Produce		json
// @Param			input	body	utill.TodoList	true	"list info"
// @Router			/api/lists [post]
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
	listId, err := h.service.TodoList.Create(userId, intup)
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

//	@Summary		Get All Lists
//	@Security		ApiKeyAuth
//	@Tags			lists
//	@Description	get all lists
//	@ID				get-all-lists
//	@Accept			json
//	@Produce		json
//
// @Router /api/lists [get]
func (h *Handler) getAllList(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		newResponseError(c, http.StatusBadRequest, err.Error())
		return
	}
	lists, err := h.service.TodoList.GetAll(userId)
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
	id, err := utill.ParamInt(c, "id")
	if err != nil {
		newResponseError(c, http.StatusBadRequest, err.Error())
		return
	}
	list, err := h.service.TodoList.GetById(userId, id)
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
	id, err := utill.ParamInt(c, "id")
	if err != nil {
		newResponseError(c, http.StatusBadRequest, err.Error())
	}
	var input todo.InputListUpdate
	c.BindJSON(&input)
	err = h.service.TodoList.Update(userId, id, input)
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
	id, err := utill.ParamInt(c, "id")
	if err != nil {
		newResponseError(c, http.StatusBadRequest, err.Error())
		return
	}
	err = h.service.TodoList.Delete(userId, id)
	if err != nil {
		newResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.AbortWithStatusJSON(statusOk, newStatusRes("ok"))
}
