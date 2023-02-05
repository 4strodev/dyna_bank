package infrastructure

import (
	"github.com/4strodev/dyna-bank/internal/app/db/postgres"
	"github.com/4strodev/dyna-bank/internal/app/features/accounts/actions"
	"github.com/4strodev/dyna-bank/internal/app/shared/http/response"
	"github.com/4strodev/dyna-bank/internal/app/shared/http/status"
	"github.com/gofiber/fiber/v2"
)


func InitRoutes(prefix string, app *fiber.App) {
	router := app.Group(prefix)

	router.Get("/", func(c *fiber.Ctx) error {
		postgresRepository := NewPostgresAccountRepository(postgres.DBPool)
		accounts, appErr := actions.GetAccounts(postgresRepository)

		if appErr != nil {
			return appErr
		}

		response := response.HttpResponse{
			Status: status.STATUS_SUCCESS,
			Data:   accounts,
		}

		return c.JSON(response)
	})
}
