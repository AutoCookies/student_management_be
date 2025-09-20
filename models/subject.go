package models

import (
	"fmt"
)

type Subject struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Semester int    `json:"semester"`
}

func (s *Subject) GetName() string {
	return s.Name
}

func (s *Subject) GetSemester() int {
	return s.Semester
}

func (s *Subject) String() string {
	return fmt.Sprintf("Subject[ID=%d, Name=%s, Semester=%d]", s.ID, s.Name, s.Semester)
}
