package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginOptions(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Status(http.StatusOK)
}
