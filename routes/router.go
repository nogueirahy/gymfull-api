package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nogueirahy/api-gymfull/controllers"
)

func ConfigRoutes(router *echo.Echo) *echo.Echo {
	router.Use(middleware.CORS())
	main := router.Group("/api/v1")

	members := main.Group("/member")
	members.GET("/", controllers.ShowAllMembers)
	members.GET("/:id", controllers.FindMember)
	members.POST("/", controllers.CreateMember)
	members.PUT("/", controllers.EditMember)
	members.DELETE("/:id", controllers.DeleteMember)

	return router
}
