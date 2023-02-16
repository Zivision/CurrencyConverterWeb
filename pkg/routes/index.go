package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func (r routes) addIndex(rg *gin.RouterGroup) {
	route := rg.Group("/")

	route.GET("/", indexMainEndpoint)
}

func indexMainEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Message": "Hello World!",
	})
} 
