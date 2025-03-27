package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tapeds/go-fiber-template/controller"
	"github.com/tapeds/go-fiber-template/middleware"
	"github.com/tapeds/go-fiber-template/service"
)

func Event(route fiber.Router, eventController controller.EventController, jwtService service.JWTService) {
	routes := route.Group("/event")

	routes.Post("add-event", middleware.Authenticate(jwtService), eventController.CreateEvent)
	routes.Get("", middleware.Authenticate(jwtService), eventController.GetAllEvent)
	routes.Get("by-id", middleware.Authenticate(jwtService), eventController.GetEventById)
	routes.Delete("", middleware.Authenticate(jwtService), eventController.Delete)
	routes.Put("", middleware.Authenticate(jwtService), eventController.Update)
}
