package models

import (
	"time"
)

type SubjectDetails struct {
	ID          int       `json:"id"`
	TimeTableID int       `json:"timetableid"`
	SubjectID   int       `json:"subjectid"`
	Date        time.Time `json:"date"`
}

func (s *SubjectDetails) getSubjectDate() time.Time {
	return s.Date
}
