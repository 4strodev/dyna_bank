package repositories

import (
	"database/sql"

	"github.com/4strodev/dyna-bank/internal/app/features/users/domain/models"
	"github.com/4strodev/dyna-bank/internal/app/shared/errors"
)

type PostgresUserRepository struct {
	pool *sql.DB
}

func NewPostgresUserRepository(pool *sql.DB) PostgresUserRepository {
	return PostgresUserRepository{
		pool: pool,
	}
}

func (self PostgresUserRepository) Find() ([]models.User, errors.ApplicationErrorInterface) {
	var users []models.User = make([]models.User, 0)

	rows, err := self.pool.Query(`SELECT id, username, password, email, phone FROM application.users`)
	if err != nil {
		return []models.User{}, errors.NewInfrastructureError("Error getting users from database").WithMetadata(err)
	}
	defer rows.Close()

	for rows.Next() {
		var user = models.User{}
		err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.Phone)
		if err != nil {
			return []models.User{}, errors.NewInfrastructureError("Error parsing user data").WithMetadata(err)
		}

		users = append(users, user)
	}

	return users, nil
}
