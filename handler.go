package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getHandler(rtp float64) gin.HandlerFunc {
	return func(c *gin.Context) {
		m := generateMultiplier(rtp)
		c.JSON(http.StatusOK, gin.H{"result": m})
	}
}
