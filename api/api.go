package api

import (
	"github.com/gin-gonic/gin"

	"github.com/asadbekGo/market_system/api/handler"
	"github.com/asadbekGo/market_system/config"
	"github.com/asadbekGo/market_system/storage"
)

func SetUpApi(r *gin.Engine, cfg *config.Config, strg storage.StorageI) {

	handler := handler.NewHandler(cfg, strg)

	r.Use(customCORSMiddleware())

	v1 := r.Group("/v1")
	v1.Use(handler.CheckPasswordMiddleware())
	v1.Use(handler.CheckPasswordMiddleware())

	// Category ...
	v1.POST("/category", handler.CreateCategory)
	v1.GET("/category/:id", handler.GetByIDCategory)
	v1.GET("/category", handler.GetListCategory)
	v1.PUT("/category/:id", handler.UpdateCategory)
	v1.DELETE("/category/:id", handler.DeleteCategory)

	//user ...
	v1.POST("/user", handler.CreateUser)
	v1.GET("/user/:id", handler.GetByIDUser)
	v1.GET("/user", handler.GetListUser)
	v1.PUT("/user/:id", handler.UpdateUser)
	v1.DELETE("/user/:id", handler.DeleteUser)

	v1.PUT("/user/refresh", handler.RefreshToken)
	v1.POST("/user/login", handler.Login)

}

func customCORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE, HEAD")
		c.Header("Access-Control-Allow-Headers", "Password, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Max-Age", "3600")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
