package main

import (
	"errors"
	"log"

	_ "github.com/joho/godotenv/autoload"
	accounts_infrastructure "github.com/4strodev/dyna-bank/internal/app/features/accounts/infrastructure"
	app_errors "github.com/4strodev/dyna-bank/internal/app/shared/errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

var app *fiber.App

func init() {
	app = fiber.New(fiber.Config{
		ServerHeader: "Fiber",
		ErrorHandler: errorHandler,
	})

	//Use recover middleware to don't stop application
	app.Use(recover.New())
	accounts_infrastructure.InitRoutes("/accounts", app)
}

func errorHandler(c *fiber.Ctx, err error) error {
	errorMessage := "Internal server error"
	statusCode := fiber.StatusInternalServerError

	var appErr app_errors.ApplicationError
	if errors.As(err, &appErr) {
		errorMessage = appErr.Error()
		statusCode = appErr.HTTPStatusCode()
	}

	var fiberErr *fiber.Error
	if errors.As(err, &fiberErr) {
		errorMessage = fiberErr.Message
		statusCode = fiberErr.Code
	}

	log.Println(err)

	return c.Status(statusCode).JSON(fiber.Map{
		"status":  "Error",
		"message": errorMessage,
	})
}
