package router

import (
	"github.com/gin-gonic/gin"

	"gin_practice/handler"
)

type Handler struct {
	Auth *handler.Auth
	Item *handler.Item
}

func NewGinRouter(
	middlewares []gin.HandlerFunc,
	handler Handler,
) *gin.Engine {
	router := gin.New()

	for idx := range middlewares {
		router.Use(middlewares[idx])
	}

	// API version group
	v1 := router.Group("/v1")
	{
		v1.POST("/login", handler.Auth.Login())
		v1.POST("/register", handler.Auth.Register())
		v1.GET("/item", handler.Item.Items())
		v1.POST("/item", handler.Item.BulkInsert())
	}

	return router
}
