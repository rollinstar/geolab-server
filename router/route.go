package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rollinstar/geolab-server/handler"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/user")

	v1.Get("/hello", handler.Hello)
	//v1.Get("/", handler.GetAllUsers)
	//v1.Get("/:id", handler.GetSingleUser)
	//v1.Post("/", handler.CreateUser)
	//v1.Put("/:id", handler.UpdateUser)
	//v1.Delete("/:id", handler.DeleteUserByID)
}
