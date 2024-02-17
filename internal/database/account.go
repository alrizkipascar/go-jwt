package database

import (
	"database/sql"
	"fmt"

	"github.com/alrizkipascar/go-jwt/internal/models"
)

func CreateAccount(acc *models.Account) error {
	s := GetDB()
	query := `insert into account  
    (email, uuid, first_name, last_name, number, activated, encrypted_password, balance, created_at)
    values
    ( $1, $2, $3, $4, $5, $6, $7, $8, $9)
    `

	_, err := s.Query(
		query,

		acc.Email,
		acc.UUID,
		acc.FirstName,
		acc.LastName,
		acc.Number,
		acc.Activated,
		acc.EncryptedPassword,
		acc.Balance,
		acc.CreatedAt)

	if err != nil {
		return err
	}

	// fmt.Printf("%v\n", resp)

	return nil
}

func UpdateAccount(*models.Account) error {
	return nil
}

func DeleteAccount(id int) error {
	s := GetDB()

	_, err := s.Query("DELETE FROM ACCOUNT WHERE ID = $1", id)
	return err
}

func GetAccountByNumber(number int) (*models.Account, error) {
	s := GetDB()

	rows, err := s.Query("SELECT * FROM ACCOUNT WHERE number = $1", number)

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		return ScanIntoAccount(rows)
	}
	return nil, fmt.Errorf("account with number %d not found", number)
}

func GetAccountByEmail(email string) (*models.Account, error) {
	s := GetDB()

	rows, err := s.Query("SELECT * FROM ACCOUNT WHERE email = $1", email)

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		return ScanIntoAccount(rows)
	}
	return nil, fmt.Errorf("account with email %v not found", email)
}

func GetAccountByID(id int) (*models.Account, error) {
	s := GetDB()

	rows, err := s.Query("SELECT * FROM ACCOUNT WHERE ID = $1", id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		return ScanIntoAccount(rows)
	}
	return nil, fmt.Errorf("account %d not found", id)
}

func GetAccounts() ([]*models.Account, error) {
	s := GetDB()

	rows, err := s.Query("SELECT * FROM ACCOUNT")
	if err != nil {
		return nil, err
	}

	accounts := []*models.Account{}

	for rows.Next() {
		account := new(models.Account)
		account, err := ScanIntoAccount(rows)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}

	return accounts, nil
}

func ScanIntoAccount(rows *sql.Rows) (*models.Account, error) {
	account := new(models.Account)
	err := rows.Scan(
		&account.ID,
		&account.Email,
		&account.UUID,
		&account.FirstName,
		&account.LastName,
		&account.Number,
		&account.Activated,
		&account.EncryptedPassword,
		&account.Balance,
		&account.CreatedAt)
	if err != nil {
		return nil, err
	}
	return account, err
}

func ActivatedAccount(id int) error {
	s := GetDB()
	query := `
	UPDATE account
	SET activated = $1
	WHERE ID = $2
	`

	_, err := s.Query(query, 1, id)

	if err != nil {
		return err
	}

	return nil
}
