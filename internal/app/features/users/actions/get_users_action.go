package actions

import (
	"github.com/4strodev/dyna-bank/internal/app/features/users/domain/models"
	"github.com/4strodev/dyna-bank/internal/app/features/users/domain/repositories"
	"github.com/4strodev/dyna-bank/internal/app/shared/errors"
)

func GetUsersAction(repository repositories.UserRepository) ([]models.User, errors.ApplicationErrorInterface) {
	return repository.Find()
}
