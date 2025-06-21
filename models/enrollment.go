package models

import "time"

type Enrollment struct {
	ID        uint          `gorm:"primaryKey;autoIncrement" json:"id"`
	CourseID  uint          `gorm:"column:course_id" json:"course_id"`
	Course    CourseDetails `gorm:"foreignKey:ID;references:CourseID" json:"course"`
	StudentID uint          `gorm:"column:student_id" json:"student_id"`
	Student   Student       `gorm:"foreignKey:ID;references:StudentID" json:"student"`
	CreatedAt time.Time     `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time     `gorm:"autoUpdateTime" json:"updated_at"`
	TenantID  uint          `gorm:"column:tenant_id" json:"-"`
	Tenant    Tenant        `gorm:"foreignKey:TenantID;references:ID" json:"-"`
}

type CourseEnrollmentRes struct {
	ID    uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Title string `json:"title"`
}

func (CourseEnrollmentRes) TableName() string {
	return "course_details"
}

type StudentEnrollmentRes struct {
	ID        uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	FirstName string  `json:"first_name"`
	LastName  *string `json:"last_name"`
	Email     string  `json:"email"`
}

func (StudentEnrollmentRes) TableName() string {
	return "students"
}

type EnrollmentResponse struct {
	ID        uint                 `gorm:"primaryKey;autoIncrement" json:"id"`
	CourseID  uint                 `gorm:"column:course_id" json:"course_id"`
	Course    CourseEnrollmentRes  `gorm:"foreignKey:ID;references:CourseID" json:"course"`
	StudentID uint                 `gorm:"column:student_id" json:"student_id"`
	Student   StudentEnrollmentRes `gorm:"foreignKey:ID;references:StudentID" json:"student"`
	CreatedAt time.Time            `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time            `gorm:"autoUpdateTime" json:"updated_at"`
}

func (EnrollmentResponse) TableName() string {
	return "enrollments"
}

type EnrolledCourseRes struct {
	ID        uint                        `gorm:"primaryKey;autoIncrement" json:"id"`
	CourseID  uint                        `gorm:"column:course_id" json:"course_id"`
	Course    CourseDetailsPublicResponse `gorm:"foreignKey:ID;references:CourseID" json:"course"`
	StudentID uint                        `gorm:"column:student_id" json:"student_id"`
	CreatedAt time.Time                   `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time                   `gorm:"autoUpdateTime" json:"updated_at"`
}

func (EnrolledCourseRes) TableName() string {
	return "enrollments"
}
