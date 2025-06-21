package student

import (
	"dashlearn/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterStudentRoutes(rg *gin.RouterGroup) {
	authGroup := rg.Group("/private/student")
	{
		authGroup.GET("/", middleware.AuthMiddleware(), GetStudents)
		authGroup.GET("/lite", middleware.AuthMiddleware(), GetStudentLite)
		authGroup.POST("/register", middleware.AuthMiddleware(), CreateStudent)
		// userGroup.POST("/upload", UploadUser)
	}

	publicGroup := rg.Group("/student", middleware.GetTenantID())
	{
		publicGroup.POST("/login", LoginStudent)
		publicGroup.POST("/register", CreateStudentPublic)
		publicGroup.GET("/details", middleware.StudentAuthMiddleware(), GetStudentDetails)
	}
}
