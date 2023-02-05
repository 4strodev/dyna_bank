package infrastructure

import (
	"database/sql"
	"log"

	"github.com/4strodev/dyna-bank/internal/app/features/accounts/domain/models"
	"github.com/4strodev/dyna-bank/internal/app/shared/errors"
)

func NewPostgresAccountRepository(pool *sql.DB) PostgresAccountRepository {
	return PostgresAccountRepository{
		pool,
	}
}

type PostgresAccountRepository struct {
	pool *sql.DB
}

func (self PostgresAccountRepository) Find() ([]models.Account, errors.ApplicationErrorInterface) {
	var accounts []models.Account = make([]models.Account, 0)

	rows, err := self.pool.Query("SELECT id, name, user_id, summary FROM application.accounts")
	if err != nil {
		return accounts, errors.NewInfrastructureError("Error getting data from database").WithMetadata(err)
	}
	defer rows.Close()

	// Parsing values
	for rows.Next() {
		account := models.Account{}
		err = rows.Scan(&account.Id, &account.Name, &account.UserID, &account.Summary)
		if err != nil {
			log.Println(err)
			return []models.Account{}, errors.NewInfrastructureError("Error parsing accounts to struct").WithMetadata(err)
		}
		accounts = append(accounts, account)
	}

	return accounts, nil
}
