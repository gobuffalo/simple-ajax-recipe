package actions

import (
	"time"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/middleware"
	"github.com/gobuffalo/envy"
	"github.com/markbates/going/defaults"

	"github.com/gobuffalo/buffalo/middleware/csrf"
)

// ENV is used to help switch settings based on where the
// application is being run. Default is "development".
var ENV = envy.Get("GO_ENV", "development")
var app *buffalo.App

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
func App() *buffalo.App {
	if app == nil {
		app = buffalo.New(buffalo.Options{
			Env:         ENV,
			SessionName: "_coke_session",
		})
		// Automatically redirect to SSL

		if ENV == "development" {
			app.Use(middleware.ParameterLogger)
		}

		// Protect against CSRF attacks. https://www.owasp.org/index.php/Cross-Site_Request_Forgery_(CSRF)
		// Remove to disable this.
		app.Use(csrf.New)
		app.Use(func(next buffalo.Handler) buffalo.Handler {
			return func(c buffalo.Context) error {
				c.Set("badge", defaults.String(c.Request().FormValue("badge"), "warning"))
				c.Set("now", time.Now())
				return next(c)
			}
		})

		app.POST("/traffic", TrafficCop)
		app.GET("/", HomeHandler).Alias("/traffic")

		app.ServeFiles("/", assetsBox) // serve files from the public directory
	}

	return app
}
