package actions

import (
	"errors"

	"github.com/gobuffalo/buffalo"
	"github.com/gopheracademy/slideshow/models"
)

// CoursesList default implementation.
func CoursesList(c buffalo.Context) error {
	cl, err := models.GetCourseList()
	if err != nil {
		return c.Error(500, err)
	}
	c.Set("courses", cl)
	return c.Render(200, r.HTML("courses/list.html"))
}

// CoursesShow default implementation.
func CoursesShow(c buffalo.Context) error {
	id, err := c.ParamInt("id")
	if err != nil {
		return c.Error(404, errors.New("Not Found"))
	}
	course, err := models.GetFullCourse(id)
	if err != nil {
		return c.Error(500, err)
	}
	c.Set("course", course)
	r.HTMLLayout = "slidemaster.html"
	return c.Render(200, r.HTML("courses/slide.html"))
}
