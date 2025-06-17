package enrollment

import (
	"dashlearn/models"
	"dashlearn/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateEnrollment(ctx *gin.Context) {
	var input models.Enrollment
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newEnrollment := models.Enrollment{
		StudentID: input.StudentID,
		CourseID:  input.CourseID,
		TenantID:  ctx.GetUint("tenant_id"),
	}

	if err := utils.DB.Create(&newEnrollment).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Enrollment created successfully"})
}

func GetEnrollments(ctx *gin.Context) {
	var enrollments []models.EnrollmentResponse

	if err := utils.DB.
		Where("tenant_id = ?", ctx.GetUint("tenant_id")).
		Preload("Student", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "first_name", "last_name", "email")
		}).
		Preload("Course", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "title")
		}).
		Find(&enrollments).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": enrollments})
}
