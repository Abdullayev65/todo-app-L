package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ResError struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json:"status"`
}

func newStatusRes(status string) *statusResponse {
	return &statusResponse{Status: status}
}

func newResponseError(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, ResError{message})
}
