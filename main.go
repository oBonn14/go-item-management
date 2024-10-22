package main

import (
	"Go-CRUD/app"
	"Go-CRUD/app/domain/customer"
	"Go-CRUD/app/domain/handler"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/joho/godotenv/autoload"
	"os"
	"strconv"
)

func main() {
	db := app.ConnectToDatabase()

	customerRepo := customer.NewCustomerRepository(db)
	customerService := customer.NewService(customerRepo)
	customerHandler := handler.NewCustomerHandler(customerService)

	api := fiber.New()
	api.Use(logger.New())

	route := api.Group("/")

	route.Get("/", homeHandler())

	route.Post("/customer", customerHandler.AddCustomer)
	route.Get("/customer", customerHandler.GetAllCustomer)
	route.Put("/customer/:id", customerHandler.UpdateCustomer)
	route.Delete("/customer/:id", customerHandler.DeleteCustomer)

	port, _ := strconv.Atoi(os.Getenv("PORT"))
	err := api.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		panic(fmt.Sprintf("Cannot start server: %s", err))
	}
}

func homeHandler() fiber.Handler {
	return func(c *fiber.Ctx) error { return c.SendString("Hello Udah Jalan") }
}
