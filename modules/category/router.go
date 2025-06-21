package category

import (
	"dashlearn/middleware"
	"dashlearn/utils"

	"github.com/gin-gonic/gin"
)

func RegisterCategoryRoutes(rg *gin.RouterGroup) {

	handler := NewCategoryHandler(utils.DB)

	authgroup := rg.Group("/private/category", middleware.AuthMiddleware())
	{
		authgroup.GET("/", handler.GetAll)
		authgroup.GET("/:id", handler.GetByID)
		authgroup.POST("/create", handler.Create)
		authgroup.PUT("/update/:id", handler.Update)
		// routerGroup.DELETE("/delete/:id", middleware.AuthMiddleware(), DeleteCategory)
	}

	publicGroup := rg.Group("/category", middleware.GetTenantID())
	{
		publicGroup.GET("/", handler.GetAll)
		// publicGroup.GET("/:id", GetCategoryByIDPublic)
	}
}
