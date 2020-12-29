package persistence

import (
	"github.com/jmoiron/sqlx"
)

type topicsPersistencePostgres struct{}

func NewTopicsPersistence() repository.TopicsRepository {
	return &topicsPersistencePostgres{}
}

type accountsAdminPersistencePostgres struct {
	db *sqlx.DB
}

func (app topicsPersistencePostgres) FetchAccount(userID int, externalID string) (*model.Account, error) {
	const selectAccountSQL = `
    SELECT
        id
      , user_id
      , provider_code
    FROM
        accounts
    `

	var account model.Account
	err := db.With(func(conn db.Conn) error {
		return conn.QueryRow(selectAccountSQL, userID, externalID).Scan(
			&account.ID,
			&account.UserID,
			&account.ProviderCode,
			&account.UID,
			&account.UName,
		)
	})

	if errors.Cause(err) == db.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &account, nil
}
