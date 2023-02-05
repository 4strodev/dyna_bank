package infrastructure

import (
	"github.com/4strodev/dyna-bank/internal/app/db/postgres"
	"github.com/4strodev/dyna-bank/internal/app/features/users/actions"
	"github.com/4strodev/dyna-bank/internal/app/features/users/infrastructure/repositories"
	"github.com/4strodev/dyna-bank/internal/app/shared/http/response"
	"github.com/4strodev/dyna-bank/internal/app/shared/http/status"
	"github.com/gofiber/fiber/v2"
)

func InitRoutes(prefix string, app *fiber.App) {
	router := app.Group(prefix)
	router.Get("/", func (c *fiber.Ctx) error {
		postgresUsersRepository := repositories.NewPostgresUserRepository(postgres.DBPool)
		users, err := actions.GetUsersAction(postgresUsersRepository)
		if err != nil {
			return err
		}

		response := response.HttpResponse{
			Status: status.STATUS_SUCCESS,
			Data: users,
		}

		return c.JSON(response)
	})
}
