package enrollment

import (
	"dashlearn/models"
	"dashlearn/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateEnrollment(c *gin.Context) {
	var input models.Enrollment
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newEnrollment := models.Enrollment{
		StudentID: input.StudentID,
		CourseID:  input.CourseID,
		TenantID:  c.GetUint("tenant_id"),
	}

	if err := utils.DB.Create(&newEnrollment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Enrollment created successfully"})
}

func GetEnrollments(c *gin.Context) {
	var enrollments []models.EnrollmentResponse

	if err := utils.DB.
		Where("tenant_id = ?", c.GetUint("tenant_id")).
		Preload("Student", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "first_name", "last_name", "email")
		}).
		Preload("Course", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "title")
		}).
		Find(&enrollments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": enrollments})
}

func GetEnrolledCourses(c *gin.Context) {
	var enrollments []models.EnrolledCourseRes

	utils.DB.
		Where(&models.Enrollment{
			TenantID:  c.GetUint("tenant_id"),
			StudentID: c.GetUint("user_id"),
		}).
		Preload("Course").
		Find(&enrollments)

	c.JSON(http.StatusOK, gin.H{"data": enrollments})
}
