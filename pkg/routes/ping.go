package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func (r routes) addPing(rg *gin.RouterGroup) {
	route := rg.Group("/ping")

	route.GET("/", pingMainEndpoint)
}

func pingMainEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Message": "Pong!",
	})
} 
