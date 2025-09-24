package models

import (
	"fmt"
	"time"
)

type TimeTable struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

// Constructor: tạo TimeTable mới với CreatedAt = time.Now()
func NewTimeTable(id int, name string) *TimeTable {
	return &TimeTable{
		ID:        id,
		Name:      name,
		CreatedAt: time.Now(),
	}
}

// Getters
func (t *TimeTable) GetID() int              { return t.ID }
func (t *TimeTable) GetName() string         { return t.Name }
func (t *TimeTable) GetCreatedAt() time.Time { return t.CreatedAt }

// Setters
func (t *TimeTable) SetID(id int)                     { t.ID = id }
func (t *TimeTable) SetName(name string)              { t.Name = name }
func (t *TimeTable) SetCreatedAt(createdAt time.Time) { t.CreatedAt = createdAt }

// String method
func (t *TimeTable) String() string {
	return fmt.Sprintf(
		"TimeTable[ID=%d, Name=%s, CreatedAt=%s]",
		t.ID,
		t.Name,
		t.CreatedAt.Format("2006-01-02 15:04:05"),
	)
}
