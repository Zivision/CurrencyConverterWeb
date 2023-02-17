package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func (r routes) addConvert(rg *gin.RouterGroup) {
	route := rg.Group("/convert")

	route.GET("/", convertMainEndpoint)
	route.GET("/help", convertHelpEndpoint)
	route.GET("/toUSD/:id/:value", convertToUSDEndpoint)

}

type Currency struct {
	ID		string `uri:"id" binding:"required"`
	Value int		 `uri:"value" binding:"required"`
}

func convertMainEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Message": "not implemented",
	})
}

func convertHelpEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Message": "not implemented",
	})
}

func convertToUSDEndpoint(c *gin.Context) {
	var currencyType Currency
	if err := c.ShouldBindUri(&currencyType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": err,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"currency": currencyType.ID,
		"amount": currencyType.Value,
	})

}

