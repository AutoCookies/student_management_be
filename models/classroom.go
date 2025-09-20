package models

import (
	"fmt"
)

type Classroom struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Total       int    `json:"total"`
	TimeTableID int    `json:"timetableId"`
}

func (c *Classroom) getName() string {
	return c.Name
}

func (c *Classroom) getTotal() int {
	return c.Total
}

func (c *Classroom) String() string {
	return fmt.Sprintf(
		"Classroom[ID=%d, Name=%s, Total=%d, TimeTableID=%d]",
		c.ID, c.Name, c.Total, c.TimeTableID,
	)
}
