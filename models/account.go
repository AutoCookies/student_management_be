package models

import (
	"fmt"
	"time"
)

type Account struct {
	ID          int       `json:"id"`
	Username    string    `json:"name"`
	Email       string    `json:"email"`
	Birth       time.Time `json:"birth"`
	Role        string    `json:"role"`
	CreatedAt   time.Time `json:"createdAt"`
	TimeTableID *int      `json:"timetableId"`
}

func NewAccount(id int, name string, email string, birth time.Time, role string, timetableId *int) *Account {
	return &Account{
		ID:          id,
		Username:    name,
		Email:       email,
		Birth:       birth,
		Role:        role,
		CreatedAt:   time.Now(),
		TimeTableID: timetableId,
	}
}

func (a *Account) GetName() string         { return a.Username }
func (a *Account) GetEmail() string        { return a.Email }
func (a *Account) GetRole() string         { return a.Role }
func (a *Account) GetBirth() time.Time     { return a.Birth }
func (a *Account) GetTimeTableID() *int    { return a.TimeTableID }
func (a *Account) GetCreatedAt() time.Time { return a.CreatedAt }

func (a *Account) SetName(name string)      { a.Username = name }
func (a *Account) SetEmail(email string)    { a.Email = email }
func (a *Account) SetRole(role string)      { a.Role = role }
func (a *Account) SetBirth(birth time.Time) { a.Birth = birth }
func (a *Account) SetTimeTableID(id *int)   { a.TimeTableID = id }

func (a *Account) String() string {
	tid := "nil"
	if a.TimeTableID != nil {
		tid = fmt.Sprintf("%d", *a.TimeTableID)
	}
	return fmt.Sprintf(
		"Account[ID=%d, Username=%s, Email=%s, Birth=%s, Role=%s, TimeTableID=%s, CreatedAt=%s]",
		a.ID,
		a.Username,
		a.Email,
		a.Birth.Format("2006-01-02"),
		a.Role,
		tid,
		a.CreatedAt.Format("2006-01-02 15:04:05"),
	)
}
