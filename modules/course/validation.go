package course

import (
	"dashlearn/models"
	"time"
)

type CourseDetailsInput struct {
	Title           string                `form:"title" json:"title" binding:"required"`
	Summary         string                `form:"summary" json:"summary" binding:"required"`
	Description     *string               `form:"description" json:"description" binding:"omitempty"`
	Visibility      models.Visibility     `form:"visibility" json:"visibility" binding:"required,oneof=public private protected"`
	IsScheduled     *bool                 `form:"is_scheduled" json:"is_scheduled" binding:"omitempty"`
	ScheduleDate    *time.Time            `form:"schedule_date" json:"schedule_date" binding:"omitempty"`
	ScheduleTime    *time.Time            `form:"schedule_time" json:"schedule_time" binding:"omitempty"`
	PricingModel    models.PricingModel   `form:"pricing_model" json:"pricing_model" binding:"omitempty"`
	RegularPrice    *float32              `form:"regular_price" json:"regular_price" binding:"omitempty"`
	SalePrice       *float32              `form:"sale_price" json:"sale_price" binding:"omitempty"`
	ShowCommingSoom *bool                 `form:"show_comming_soom" json:"show_comming_soom" binding:"omitempty"`
	Tags            *[]string             `form:"tags" json:"tags" binding:"omitempty"`
	AuthorID        uint                  `form:"author_id" json:"author_id" binding:"required"`
	FeaturedImage   *string               `form:"featured_image" json:"featured_image" binding:"omitempty"`
	IntroVideo      *models.IntroVideo    `form:"intro_video" json:"intro_video" binding:"omitempty"`
	Overview        *[]string             `form:"overview" json:"overview" binding:"omitempty"`
	CourseChapters  []CreateCourseChapter `form:"course_chapters" json:"course_chapters"`
	GeneralSettings CreateGeneralSettings `form:"general_settings" json:"general_settings" binding:"required"`
	Instructors     []int32               `form:"course_instructors" json:"course_instructors"`
}

type CreateCourseDetailsInput struct {
	Title           string              `form:"title" json:"title" binding:"required"`
	Summary         string              `form:"summary" json:"summary" binding:"required"`
	Description     *string             `form:"description" json:"description" binding:"omitempty"`
	Visibility      models.Visibility   `form:"visibility" json:"visibility" binding:"required,oneof=public private protected"`
	IsScheduled     *bool               `form:"is_scheduled" json:"is_scheduled" binding:"omitempty"`
	ScheduleDate    *time.Time          `form:"schedule_date" json:"schedule_date" binding:"omitempty"`
	ScheduleTime    *time.Time          `form:"schedule_time" json:"schedule_time" binding:"omitempty"`
	PricingModel    models.PricingModel `form:"pricing_model" json:"pricing_model" binding:"omitempty"`
	RegularPrice    *float32            `form:"regular_price" json:"regular_price" binding:"omitempty"`
	SalePrice       *float32            `form:"sale_price" json:"sale_price" binding:"omitempty"`
	ShowCommingSoom *bool               `form:"show_comming_soom" json:"show_comming_soom" binding:"omitempty"`
	Tags            *[]string           `form:"tags" json:"tags" binding:"omitempty"`
	AuthorID        uint                `form:"author_id" json:"author_id" binding:"required"`
	IntroVideo      *models.IntroVideo  `form:"intro_video" json:"intro_video" binding:"omitempty"`
	Overview        *[]string           `form:"overview" json:"overview" binding:"omitempty"`
	// FeaturedImage   *string             `form:"featured_image" binding:"omitempty"`
}

type CreateCourseChapter struct {
	ID            *int64               `json:"id" form:"id" binding:"omitempty"`
	Position      int32                `json:"position" form:"position" binding:"required"`
	Title         string               `json:"title" form:"title" binding:"required"`
	Description   *string              `json:"description" form:"description" binding:"omitempty"`
	Access        models.Access        `json:"access" form:"access" binding:"required,oneof=draft published"`
	CourseLessons []CreateCourseLesson `form:"course_lessons" json:"course_lessons"`
}

type CreateCourseLesson struct {
	ID          *int64                  `json:"id" form:"id" binding:"omitempty"`
	Title       string                  `json:"title" form:"title" binding:"required"`
	Description *string                 `json:"description" form:"description" binding:"omitempty"`
	LessonType  models.LessonType       `json:"lesson_type" form:"lesson_type" binding:"required,oneof=video live_session audio text"`
	SourceType  models.LessonSourceType `json:"source_type" form:"source_type" binding:"required,oneof=youtube vimeo sound_cloud spotify custom_code upload"`
	Source      models.Source           `json:"source" form:"source" binding:"omitempty"`
	IsPublished bool                    `json:"is_published" form:"is_published" binding:"required"`
	IsPublic    bool                    `json:"is_public" form:"is_public" binding:"required"`
	Resources   *[]any                  `json:"resources" form:"resources" binding:"omitempty"`
}

type CreateGeneralSettings struct {
	DifficultyLevel models.DifficultyLevel `json:"difficulty_level" form:"difficulty_level" binding:"required,oneof=all beginner intermediate expert"`
	MaximumStudent  *int32                 `json:"maximum_student" form:"maximum_student" binding:"omitempty"`
	Language        *string                `json:"language" form:"language" binding:"omitempty"`
	CategoryID      uint                   `json:"category_id" form:"category_id" binding:"required"`
	Duration        *string                `json:"duration" form:"duration" binding:"omitempty"`
}
