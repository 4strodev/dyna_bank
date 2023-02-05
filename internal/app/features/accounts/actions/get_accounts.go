package actions

import (
	"github.com/4strodev/dyna-bank/internal/app/features/accounts/domain/models"
	"github.com/4strodev/dyna-bank/internal/app/features/accounts/domain/repositories"
	"github.com/4strodev/dyna-bank/internal/app/shared/errors"
)

func GetAccounts(repository repositories.AccountRepository) ([]models.Account, errors.ApplicationErrorInterface) {
	return repository.Find()
}
