package course

import (
	"dashlearn/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterCourseRoutes(rg *gin.RouterGroup) {
	routerGroup := rg.Group("/course")
	{
		routerGroup.POST("/create", middleware.AuthMiddleware(), CreateCourse)
	}
}
