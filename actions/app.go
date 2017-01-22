package actions

import (
	"os"

	"github.com/gobuffalo/buffalo"
	"github.com/markbates/going/defaults"
)

// ENV is used to help switch settings based on where the
// application is being run. Default is "development".
var ENV = defaults.String(os.Getenv("GO_ENV"), "development")
var app *buffalo.App

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
func App() *buffalo.App {
	if app == nil {
		app = buffalo.Automatic(buffalo.Options{
			Env: ENV,
		})

		//	app.Use(middleware.PopTransaction(models.DB))

		app.GET("/", HomeHandler)

		app.GET("/courses", CoursesList)
		app.GET("/courses/{id}", CoursesShow)
		app.ServeFiles("/assets", assetsPath())
	}

	return app
}
