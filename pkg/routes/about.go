package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func (r routes) addAbout(rg *gin.RouterGroup) {
	route := rg.Group("/about")

	route.GET("/", aboutMainEndpoint)
}

func aboutMainEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Title": "about page",
	})
} 
