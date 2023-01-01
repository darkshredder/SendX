package routes

import (
	"sendx/handlers"
	"sendx/utils"

	"github.com/gofiber/fiber/v2"
)

// Register is the main route group
func Register(app *fiber.App, jobs chan utils.Job) {

	// Serve static files
	app.Static("/files", "./downloads", fiber.Static{
		Compress: false,
	})

	// Page Source Route file (routes/pageSource.go)
	app.Post("/pagesource", func(c *fiber.Ctx) error {
		return handlers.GetPageSource(c, jobs)
	})

	app.Get("/*", func(c *fiber.Ctx) error {
		return c.SendFile("./downloads/not_found.html")
	})

}
