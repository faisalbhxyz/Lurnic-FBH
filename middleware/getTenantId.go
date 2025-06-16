package middleware

import (
	"dashlearn/models"
	"dashlearn/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTenantID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		appKeyHeader := ctx.GetHeader("app-key")

		if appKeyHeader == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "App key (app-key) header is missing"})
			ctx.Abort()
			return
		}

		// get tenantID
		var tenant models.Tenant
		utils.DB.Where("app_key = ?", appKeyHeader).Select("id", "app_key").First(&tenant)

		if tenant.ID == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Tenant not found"})
			ctx.Abort()
			return
		}

		ctx.Set("tenant_id", tenant.ID)

		ctx.Next()
	}
}
