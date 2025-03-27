package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tapeds/go-fiber-template/controller"
	"github.com/tapeds/go-fiber-template/middleware"
	"github.com/tapeds/go-fiber-template/service"
)

func Transaction(route fiber.Router, transactionController controller.TransactionController, jwtService service.JWTService) {
	routes := route.Group("/transaction")

	routes.Post("add-transaction", middleware.Authenticate(jwtService), transactionController.CreateTransaction)
	routes.Get("", middleware.Authenticate(jwtService), transactionController.GetAllTransactions)
	routes.Get("by-id", middleware.Authenticate(jwtService), transactionController.GetTransactionById)
	routes.Delete(":id", middleware.Authenticate(jwtService), transactionController.DeleteTransaction)
	routes.Put(":id", middleware.Authenticate(jwtService), transactionController.UpdateTransaction)
}
