package repository

import (
	"database/sql"
	"errors"
	"log"
	"student_management_be/models"
	"time"
)

// MySQLAccountRepository implements AccountRepository using MySQL
type MySQLAccountRepository struct {
	DB *sql.DB
}

func NewMySQLAccountRepository(db *sql.DB) *MySQLAccountRepository {
	return &MySQLAccountRepository{DB: db}
}

// Add inserts a new account into MySQL
func (r *MySQLAccountRepository) Add(account *models.Account) error {
	query := `
	INSERT INTO Account (Username, Email, Birth, Role, TimeTableID, CreatedAt)
	VALUES (?, ?, ?, ?, ?, ?)
	`
	result, err := r.DB.Exec(query,
		account.Username,
		account.Email,
		account.Birth,
		account.Role,
		account.TimeTableID, // *int sẽ tự chuyển NULL nếu nil
		time.Now(),
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	account.ID = int(id)
	return nil
}

// Get returns an account by ID
func (r *MySQLAccountRepository) Get(id int) (*models.Account, error) {
	query := `
	SELECT ID, Username, Email, Birth, Role, TimeTableID, CreatedAt
	FROM Account
	WHERE ID = ?
	`
	var a models.Account
	var timetableID sql.NullInt64

	err := r.DB.QueryRow(query, id).Scan(
		&a.ID,
		&a.Username,
		&a.Email,
		&a.Birth,
		&a.Role,
		&timetableID,
		&a.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("account not found")
		}
		return nil, err
	}

	if timetableID.Valid {
		v := int(timetableID.Int64)
		a.TimeTableID = &v
	} else {
		a.TimeTableID = nil
	}

	return &a, nil
}

// Update updates an account
func (r *MySQLAccountRepository) Update(account *models.Account) error {
	query := `
	UPDATE Account
	SET Username=?, Email=?, Birth=?, Role=?, TimeTableID=?
	WHERE ID=?
	`
	result, err := r.DB.Exec(query,
		account.Username,
		account.Email,
		account.Birth,
		account.Role,
		account.TimeTableID, // *int nil sẽ chuyển NULL
		account.ID,
	)
	if err != nil {
		return err
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("account not found")
	}
	return nil
}

// Delete deletes an account by ID
func (r *MySQLAccountRepository) Delete(id int) error {
	result, err := r.DB.Exec("DELETE FROM Account WHERE ID=?", id)
	if err != nil {
		return err
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("account not found")
	}
	return nil
}

// List returns all accounts
func (r *MySQLAccountRepository) List() ([]*models.Account, error) {
	query := `
	SELECT ID, Username, Email, Birth, Role, TimeTableID, CreatedAt
	FROM Account
	`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var accounts []*models.Account
	for rows.Next() {
		var a models.Account
		var timetableID sql.NullInt64
		if err := rows.Scan(&a.ID, &a.Username, &a.Email, &a.Birth, &a.Role, &timetableID, &a.CreatedAt); err != nil {
			log.Println("Scan error:", err)
			continue
		}
		if timetableID.Valid {
			v := int(timetableID.Int64)
			a.TimeTableID = &v
		} else {
			a.TimeTableID = nil
		}
		accounts = append(accounts, &a)
	}
	return accounts, nil
}
