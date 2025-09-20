package models

import (
	"fmt"
	"time"
)

type Account struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Birth       time.Time `json:"birth"`
	Role        string    `json:"role"`
	TimeTableID int       `json:"timetableId"`
}

func (a *Account) GetName() string {
	return a.Name
}

func (a *Account) GetEmail() string {
	return a.Email
}

func (a *Account) GetRole() string {
	return a.Role
}

func (a *Account) GetBirth() time.Time {
	return a.Birth
}

func (a *Account) String() string {
	return fmt.Sprintf(
		"Account[ID=%d, Name=%s, Email=%s, Birth=%s, Role=%s, TimeTableID=%d]",
		a.ID,
		a.Name,
		a.Email,
		a.Birth.Format("2006-01-02"),
		a.Role,
		a.TimeTableID,
	)
}
