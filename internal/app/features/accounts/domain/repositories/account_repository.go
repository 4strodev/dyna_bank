package repositories

import (
	account_models "github.com/4strodev/dyna-bank/internal/app/features/accounts/domain/models"
	"github.com/4strodev/dyna-bank/internal/app/shared/errors"
)

type AccountRepository interface {
	Find() ([]account_models.Account, errors.ApplicationError)
}
