package instructor

import (
	"dashlearn/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterInstructorRoutes(rg *gin.RouterGroup) {
	routesGroup := rg.Group("/instructor")
	{
		routesGroup.GET("/", middleware.AuthMiddleware(), GetInstructors)
		routesGroup.POST("/register", middleware.AuthMiddleware(), CreateInstructor)
		// userGroup.POST("/login", LoginUser)
		// userGroup.POST("/upload", UploadUser)
	}
}
