package routes

import (
	"github.com/gin-gonic/gin"
)

type routes struct {
	router *gin.Engine
}

func NewRoutes() routes {
	r := routes{
		router: gin.Default(),
	}
	root := r.router.Group("/")

	r.addIndex(root)
	r.addAbout(root)
	r.addPing(root)

	return r
}

func (r routes) Run(addr ...string) error {
	return r.router.Run()
}