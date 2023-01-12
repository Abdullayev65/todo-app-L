package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userIdCtx           = "userId"
	statusOk            = http.StatusOK
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newResponseError(c, http.StatusUnauthorized, "Authorization header is empty")
		return
	}
	fields := strings.Fields(header)
	if len(fields) != 2 {
		newResponseError(c, http.StatusUnauthorized, "Authorization header is invalid")
		return
	}
	userId, err := h.service.ParseToken(fields[1])
	if err != nil {
		newResponseError(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.Set(userIdCtx, userId)
}

func (h Handler) getUserId(c *gin.Context) (int, error) {
	userId, ok := c.Get(userIdCtx)
	if !ok {
		return 0, errors.New("userId not founded")
	}
	return userId.(int), nil
}
