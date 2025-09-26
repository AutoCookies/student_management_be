package repository

import (
	"database/sql"
	"errors"
	"student_management_be/models"
	"time"
)

type MySQLClassroomRepository struct {
	DB *sql.DB
}

func NewMySQLClassroomRepository(db *sql.DB) *MySQLClassroomRepository {
	return &MySQLClassroomRepository{DB: db}
}

func (r *MySQLClassroomRepository) Add(classroom *models.Classroom) error {
	query := `
		INSERT INTO classroom (Name, Total, TimeTableID, CreatedAt)
		VALUES (?, ?, ?, ?)
	`

	result, err := r.DB.Exec(query,
		classroom.Name,
		classroom.Total,
		classroom.TimeTableID,
		time.Now(),
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	classroom.ID = int(id)
	return nil
}

func (r *MySQLClassroomRepository) Get(id int) (*models.Classroom, error) {
	query := `
		SELECT ID, Name, Total, TimeTableID, CreatedAt
		FROM classroom
		WHERE ID = ?
	`

	var a models.Classroom
	var timetableID sql.NullInt64

	err := r.DB.QueryRow(query, id).Scan(
		&a.ID,
		&a.Name,
		&a.Total,
		&timetableID,
		&a.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("classroom not found")
		}
		return nil, err
	}

	if timetableID.Valid {
		val := int(timetableID.Int64)
		a.TimeTableID = &val
	} else {
		a.TimeTableID = nil
	}

	return &a, nil
}

func (r *MySQLClassroomRepository) Update(classroom *models.Classroom) error {
	query := `
		UPDATE classroom 
		SET Name = ?, Total = ?, TimeTableID = ? 
		WHERE ID = ?
	`

	result, err := r.DB.Exec(query,
		classroom.Name,
		classroom.Total,
		classroom.TimeTableID,
		classroom.ID,
	)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("classroom not found")
	}

	return nil
}

func (r *MySQLClassroomRepository) Delete(id int) error {
	query := `DELETE FROM classroom WHERE ID = ?`

	result, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return errors.New("classroom not found")
	}

	return nil
}

func (r *MySQLClassroomRepository) List(limit, offset int) ([]*models.Classroom, error) {
	query := `
		SELECT ID, Name, Total, TimeTableID, CreatedAt
		FROM classroom
		ORDER BY CreatedAt DESC
		LIMIT ? OFFSET ?
	`

	rows, err := r.DB.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var classrooms []*models.Classroom
	for rows.Next() {
		var c models.Classroom
		var timetableID sql.NullInt64

		if err := rows.Scan(
			&c.ID,
			&c.Name,
			&c.Total,
			&timetableID,
			&c.CreatedAt,
		); err != nil {
			return nil, err
		}

		if timetableID.Valid {
			val := int(timetableID.Int64)
			c.TimeTableID = &val
		} else {
			c.TimeTableID = nil
		}

		classrooms = append(classrooms, &c)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return classrooms, nil
}
