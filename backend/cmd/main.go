package main

import (
	"github.com/gin-gonic/gin"

	"github.com/sdkim96/mcp-marketplace/internal/config"
	"github.com/sdkim96/mcp-marketplace/internal/db"
	"github.com/sdkim96/mcp-marketplace/internal/middleware"
	"github.com/sdkim96/mcp-marketplace/internal/routers"
)

func main() {

	engine := gin.Default()
	appConfig := config.GetAppConfig()
	db.InitDB()

	err := GroupRouter(engine, appConfig)
	if err != nil {
		panic(err)
	}
	engine.Run(":8080")
}

// Grouping routes
func GroupRouter(r *gin.Engine, c *config.AppConfig) error {

	r.Use(middleware.GlobalApplicationMiddleware(c))

	public := r.Group("/api/v1")
	public.GET("/health", routers.HealthCheck)
	public.POST("/login", routers.Login)
	public.POST("/sign-up", routers.Signup)

	me := r.Group("/api/v1/me")
	me.Use(middleware.AuthMiddleware())
	me.GET("", routers.GetMe)

	return nil
}
