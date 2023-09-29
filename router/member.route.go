package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rollinstar/geolab-server/handler"
)

func SetupMemberRoutes(app *fiber.App) {
	api := app.Group("/geolab/api/v1")
	v1 := api.Group("/member")
	v1.Get("/user/:id", handler.GetUser)
	v1.Post("/user", handler.CreateUser)
}
