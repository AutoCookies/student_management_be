package models

import (
	"fmt"
	"time"
)

type Test struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	AccountId int       `json:"accountId"`
}

// Constructor: tạo Test mới với CreatedAt = time.Now()
func NewTest(id int, accountId int) *Test {
	return &Test{
		ID:        id,
		CreatedAt: time.Now(),
		AccountId: accountId,
	}
}

// Getters
func (t *Test) GetID() int              { return t.ID }
func (t *Test) GetCreatedAt() time.Time { return t.CreatedAt }
func (t *Test) GetAccountId() int       { return t.AccountId }

// Setters
func (t *Test) SetID(id int)                     { t.ID = id }
func (t *Test) SetCreatedAt(createdAt time.Time) { t.CreatedAt = createdAt }
func (t *Test) SetAccountId(accountId int)       { t.AccountId = accountId }

// String method
func (t *Test) String() string {
	return fmt.Sprintf(
		"Test[ID=%d, CreatedAt=%s, AccountId=%d]",
		t.ID,
		t.CreatedAt.Format("2006-01-02 15:04:05"),
		t.AccountId,
	)
}
