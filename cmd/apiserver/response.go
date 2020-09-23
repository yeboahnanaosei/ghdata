package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type payload struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data,omitempty"`
}

func routeNotFound(c *gin.Context) {
	c.JSON(http.StatusOK, payload{
		Success: false,
		Code: http.StatusNotFound,
		Msg: "Resource not found",
	})
}
