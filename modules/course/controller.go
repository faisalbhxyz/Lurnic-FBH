package course

import (
	"context"
	"dashlearn/models"
	"dashlearn/utils"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jinzhu/copier"
)

func GetCourses(c *gin.Context) {
	var courses []models.CourseDetailsResponse

	if err := utils.DB.Where("tenant_id = ?", c.GetUint("tenant_id")).Preload("Author").Preload("Chapters").Preload("Chapters.Lessons").Preload("GeneralSettings").Preload("GeneralSettings.Category").Preload("Instructors").Preload("Instructors.Instructor").Find(&courses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": courses})
}

func GetCoursesLite(c *gin.Context) {
	var courses []struct {
		ID    uint   `json:"id"`
		Title string `json:"title"`
	}

	if err := utils.DB.Table("course_details").Where("tenant_id = ?", c.GetUint("tenant_id")).Select("id", "title", "tenant_id").Find(&courses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": courses})
}

func GetPublicCourses(c *gin.Context) {
	var courses []models.CourseDetailsPublicResponse

	if err := utils.DB.Where("tenant_id = ?", c.GetUint("tenant_id")).Preload("GeneralSettings").Preload("GeneralSettings.Category").Find(&courses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": courses})
}

func GetCourseByID(c *gin.Context) {
	courseID := c.Param("id")

	var course models.CourseDetailsResponse

	if err := utils.DB.Model(&models.CourseDetails{}).
		Where("tenant_id = ? AND id = ?", c.GetUint("tenant_id"), courseID).
		Preload("Author").
		Preload("Chapters").
		Preload("Chapters.Lessons").
		Preload("GeneralSettings").
		Preload("GeneralSettings.Category").
		Preload("Instructors").
		Preload("Instructors.Instructor").
		Preload("Enrollments").
		First(&course).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": course})
}

func CreateCourse(c *gin.Context) {
	var input CourseDetailsInput
	var flatInput CreateCourseDetailsInput

	// Step 1: Bind all flat fields (this ignores nested JSON fields like course_chapters)
	if err := c.ShouldBindWith(&flatInput, binding.FormMultipart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_ = copier.Copy(&input, &flatInput)

	// Step 2: Manually parse nested JSON fields from string values
	if chaptersStr := c.PostForm("course_chapters"); chaptersStr != "" {
		var courseChapters []CreateCourseChapter
		if err := json.Unmarshal([]byte(chaptersStr), &courseChapters); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course_chapters: " + err.Error()})
			return
		}
		input.CourseChapters = courseChapters
	}

	if generalSettingsStr := c.PostForm("general_settings"); generalSettingsStr != "" {
		var generalSettings CreateGeneralSettings
		if err := json.Unmarshal([]byte(generalSettingsStr), &generalSettings); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid general_settings: " + err.Error()})
			return
		}
		input.GeneralSettings = generalSettings
	}

	if instructorsStr := c.PostForm("course_instructors"); instructorsStr != "" {
		var instructors []int32
		if err := json.Unmarshal([]byte(instructorsStr), &instructors); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid instructors: " + err.Error()})
			return
		}
		input.Instructors = instructors
	}

	file, err := c.FormFile("featured_image")
	if err == nil {
		url, err := utils.UploadFile(context.Background(), file)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		input.FeaturedImage = &url
	} else {
		input.FeaturedImage = nil
	}

	// // Step 3: Debug log the final parsed object
	// if output, err := json.MarshalIndent(input, "", "  "); err == nil {
	// 	fmt.Println("Parsed Input:\n", string(output))
	// }

	//create course details

	var videoPtr *models.IntroVideo

	if input.IntroVideo == nil ||
		(input.IntroVideo.Type == "" && input.IntroVideo.Source == "") {
		// Set pointer to nil to store NULL in DB
		videoPtr = nil
	} else {
		// Valid data, assign normally
		videoPtr = &models.IntroVideo{
			Type:   input.IntroVideo.Type,
			Source: input.IntroVideo.Source,
		}
	}

	input.IntroVideo = videoPtr

	tagsJSON, err := utils.NormalizeTags(input.Tags)
	if err != nil {
		tagsJSON = nil
	}

	overviewJSON, err := utils.NormalizeTags(input.Overview)
	if err != nil {
		overviewJSON = nil
	}

	var introVideo *utils.JSONB[models.IntroVideo]
	if input.IntroVideo != nil {
		introVideo = &utils.JSONB[models.IntroVideo]{Data: *input.IntroVideo}
	} else {
		introVideo = nil
	}

	newCourseDetails := models.CourseDetails{
		Title:           input.Title,
		Summary:         input.Summary,
		Description:     utils.ZeroToNil(input.Description),
		Visibility:      input.Visibility,
		IsScheduled:     utils.ZeroToNil(input.IsScheduled),
		ScheduleDate:    utils.ZeroToNil(input.ScheduleDate),
		ScheduleTime:    utils.ZeroToNil(input.ScheduleTime),
		PricingModel:    input.PricingModel,
		RegularPrice:    input.RegularPrice,
		SalePrice:       input.SalePrice,
		ShowCommingSoom: input.ShowCommingSoom,
		Tags:            tagsJSON,
		Overview:        overviewJSON,
		FeaturedImage:   input.FeaturedImage,
		IntroVideo:      introVideo,
		AuthorID:        c.GetUint("user_id"),
		TenantID:        c.GetUint("tenant_id"),
	}

	if err := utils.DB.Create(&newCourseDetails).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// craete course chapter
	for idx, chapter := range input.CourseChapters {
		newCourseChapter := models.CourseChapter{
			CourseID:    newCourseDetails.ID,
			Title:       chapter.Title,
			Description: utils.EmptyStringToNil(chapter.Description),
			Position:    idx,
			Access:      chapter.Access,
		}

		if err := utils.DB.Create(&newCourseChapter).Error; err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if len(chapter.CourseLessons) > 0 {
			for idx, lesson := range chapter.CourseLessons {

				// sourceJSON, err := json.Marshal(lesson.Source)
				// if err != nil {
				// 	sourceJSON = nil
				// }

				sourceJSON := utils.JSONB[models.Source]{Data: lesson.Source}

				newCourseLesson := models.CourseLesson{
					ChapterID:   newCourseChapter.ID,
					Title:       lesson.Title,
					Description: utils.EmptyStringToNil(lesson.Description),
					Position:    idx,
					LessonType:  lesson.LessonType,
					SourceType:  lesson.SourceType,
					Source:      sourceJSON,
					IsPublished: lesson.IsPublished,
					IsPublic:    lesson.IsPublic,
				}
				if err := utils.DB.Create(&newCourseLesson).Error; err != nil {
					c.JSON(400, gin.H{"error": err.Error()})
					return
				}

			}
		}
	}

	// course instructors
	for _, instructor := range input.Instructors {
		newCourseInstructor := models.CourseInstructor{
			CourseID:     newCourseDetails.ID,
			InstructorID: uint(instructor),
		}
		if err := utils.DB.Create(&newCourseInstructor).Error; err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
	}

	// general settings
	var difficultyLevelPtr *models.DifficultyLevel
	if input.GeneralSettings.DifficultyLevel != "" {
		difficultyLevelPtr = &input.GeneralSettings.DifficultyLevel
	} else {
		defaultVal := models.All
		difficultyLevelPtr = &defaultVal
	}

	deafultLng := "english"

	newGeneralSettings := models.CourseGeneralSettings{
		CourseID:        newCourseDetails.ID,
		DifficultyLevel: difficultyLevelPtr,
		MaximumStudent:  utils.ZeroToNil(input.GeneralSettings.MaximumStudent),
		Language:        &deafultLng,
		CategoryID:      input.GeneralSettings.CategoryID,
		Duration:        utils.ZeroToNil(input.GeneralSettings.Duration),
	}

	if err := utils.DB.Create(&newGeneralSettings).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Course created successfully."})
}
