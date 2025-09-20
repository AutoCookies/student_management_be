package models

import "fmt"

type TimeTable struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (t *TimeTable) GetName() string {
	return t.Name
}

func (t *TimeTable) String() string {
	return fmt.Sprintf("TimeTable[ID=%d, Name=%s]", t.ID, t.Name)
}
