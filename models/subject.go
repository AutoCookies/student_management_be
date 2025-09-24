package models

import (
	"fmt"
	"time"
)

type Subject struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Semester  int       `json:"semester"`
	CreatedAt time.Time `json:"createdAt"`
}

// Constructor: tạo Subject mới với CreatedAt = time.Now()
func NewSubject(id int, name string, semester int) *Subject {
	return &Subject{
		ID:        id,
		Name:      name,
		Semester:  semester,
		CreatedAt: time.Now(),
	}
}

// Getters
func (s *Subject) GetID() int              { return s.ID }
func (s *Subject) GetName() string         { return s.Name }
func (s *Subject) GetSemester() int        { return s.Semester }
func (s *Subject) GetCreatedAt() time.Time { return s.CreatedAt }

// Setters
func (s *Subject) SetID(id int)             { s.ID = id }
func (s *Subject) SetName(name string)      { s.Name = name }
func (s *Subject) SetSemester(semester int) { s.Semester = semester }
func (s *Subject) SetCreatedAt(t time.Time) { s.CreatedAt = t }

// String method
func (s *Subject) String() string {
	return fmt.Sprintf(
		"Subject[ID=%d, Name=%s, Semester=%d, CreatedAt=%s]",
		s.ID,
		s.Name,
		s.Semester,
		s.CreatedAt.Format("2006-01-02 15:04:05"),
	)
}
