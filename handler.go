package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func testPing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
