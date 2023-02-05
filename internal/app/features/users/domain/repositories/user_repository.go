package repositories

import (
	"github.com/4strodev/dyna-bank/internal/app/features/users/domain/models"
	"github.com/4strodev/dyna-bank/internal/app/shared/errors"
)

type UserRepository interface {
	Find() ([]models.User, errors.ApplicationErrorInterface)
}
