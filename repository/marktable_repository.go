package repository

import (
	"database/sql"
	"student_management_be/models"
	"time"
)

type MySQLMarkTableRepository struct {
	DB *sql.DB
}

func NewMySQLMarkTableRepository(db *sql.DB) *MySQLMarkTableRepository {
	return &MySQLMarkTableRepository{DB: db}
}

func (r *MySQLMarkTableRepository) Add(marktable *models.MarkTable) error {
	query := `INSERT INTO MARKTABLE (CreateAt, AverageScore, Ranked, AccountId) VALUES (?, ?, ?, ?)`

	result, err := r.DB.Exec(query,
		time.Now(),
		marktable.AverageScore,
		marktable.Ranked,
		marktable.AccountId,
	)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	marktable.ID = int(id)
	return nil
}

func (r *MySQLMarkTableRepository) List(limit, offset int) ([]*models.MarkTable, error) {
	query := `SELECT * FROM MARKTABLE ORDER BY CreatedAt LIMIT ? OFFSET ?`

	rows, err := r.DB.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var marktables []*models.MarkTable
	for rows.Next() {
		var m models.MarkTable;
		var accountId sql.NullInt64

		if err := row.Scan() {
			&m.ID,
			&m.CreatedAt,
			&m.AverageScore,
			&m.AccountId
		} ; err !- nil {
			return nil, err
		}

		if 
	}
}
