package enrollment

import (
	"dashlearn/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterEnrollmentRoutes(rg *gin.RouterGroup) {

	authGroup := rg.Group("/private/enrollment", middleware.AuthMiddleware())
	{
		authGroup.GET("/", GetEnrollments)
		authGroup.POST("/create", CreateEnrollment)
	}

	publicGroup := rg.Group("/enrolled", middleware.GetTenantID())
	{
		publicGroup.GET("/courses", middleware.StudentAuthMiddleware(), GetEnrolledCourses)
	}
}
