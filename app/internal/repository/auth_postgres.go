package repository

import (
	"database/sql"
	"fmt"
	"github.com/linqcod/course-work-5/app/internal/model"
	"github.com/linqcod/course-work-5/app/pkg/postgresql"
)

type AuthPostgres struct {
	db *sql.DB
}

func NewAuthPostgres(db *sql.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user model.User) (int32, error) {
	var id int32
	query := fmt.Sprintf("INSERT INTO %s (username, email, password_hash) VALUES ($1, $2, $3) RETURNING id", postgresql.UsersTable)

	row := r.db.QueryRow(query, user.Username, user.Email, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (model.User, error) {
	var user model.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE username=$1 AND password_hash=$2", postgresql.UsersTable)

	row := r.db.QueryRow(query, username, password)

	if err := row.Scan(&user.Id, &user.Username, &user.Email, &user.Password); err != nil {
		return user, err
	}

	return user, nil
}
