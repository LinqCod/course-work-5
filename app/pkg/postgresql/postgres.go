package postgresql

import (
	"database/sql"
	"fmt"
)

const (
	UsersTable       = "users"
	DepositsTable    = "deposits"
	UserDepositTable = "users_deposits"
	BondsTable       = "bonds"
	UserBondTable    = "users_bonds"
	SharesTable      = "shares"
	UserShareTable   = "users_shares"
	CompaniesTable   = "companies"
	UserCompanyTable = "users_companies"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sql.DB, error) {
	url := fmt.Sprintf("postgres://%s:%s@localhost:%s/%s?sslmode=%s", cfg.Username, cfg.Password, cfg.Port, cfg.DBName, cfg.SSLMode)
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	//check connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
