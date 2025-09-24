package models

import (
	"fmt"
	"time"
)

type Classroom struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Total       int       `json:"total"`
	TimeTableID int       `json:"timetableId"`
	CreatedAt   time.Time `json:"createdAt"`
}

func NewClassroom(id int, name string, total int, timetableid int) *Classroom {
	return &Classroom{
		ID:          id,
		Name:        name,
		Total:       total,
		TimeTableID: timetableid,
		CreatedAt:   time.Now(),
	}
}

func (c *Classroom) GetName() string         { return c.Name }
func (c *Classroom) GetTotal() int           { return c.Total }
func (c *Classroom) GetTimeTableID() int     { return c.TimeTableID }
func (c *Classroom) GetCreatedAt() time.Time { return c.CreatedAt }

func (c *Classroom) SetName(name string)      { c.Name = name }
func (c *Classroom) SetTotal(total int)       { c.Total = total }
func (c *Classroom) SetTimeTableID(id int)    { c.TimeTableID = id }
func (c *Classroom) SetCreatedAt(t time.Time) { c.CreatedAt = t }

func (c *Classroom) String() string {
	return fmt.Sprintf(
		"Classroom[ID=%d, Name=%s, Total=%d, TimeTableID=%d]",
		c.ID, c.Name, c.Total, c.TimeTableID,
	)
}
