package utill

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func ParamInt(c *gin.Context, parName string) (int, error) {
	param := c.Param(parName)
	if param == "" {
		return 0, fmt.Errorf("param [%s] not founded", parName)
	}
	paramInt, err := strconv.Atoi(param)
	if err != nil {
		return 0, fmt.Errorf("param [%s] invalid", parName)
	}
	return paramInt, nil
}

type TodoList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}

type ListsItem struct {
	Id     int
	ListId int
	ItemId int
}
