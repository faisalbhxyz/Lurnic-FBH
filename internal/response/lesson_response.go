package response

import (
	"dashlearn/internal/models"
	"dashlearn/internal/utils"
	"time"

	"gorm.io/datatypes"
)

type CourseLessonResponse struct {
	ID          uint                       `json:"id"`
	Title       string                     `json:"title"`
	Description *string                    `json:"description"`
	LessonType  models.LessonType          `json:"lesson_type"`
	SourceType  models.LessonSourceType    `json:"source_type"`
	Source      utils.JSONB[models.Source] `json:"source"`
	IsPublished bool                       `json:"is_published,omitempty"`
	IsPublic    bool                       `json:"is_public"`
	Resources   datatypes.JSON             `json:"resources,omitempty"` // filename, mimetype, url, size
	Position    int                        `json:"position,omitempty"`
	CreatedAt   time.Time                  `json:"created_at"`
	UpdatedAt   time.Time                  `json:"updated_at,omitempty"`
	ChapterID   uint                       `json:"chapter_id"`
}
