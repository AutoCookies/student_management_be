package models

import (
	"fmt"
	"time"
)

type MarkTable struct {
	ID           int       `json:"id"`
	CreatedAt    time.Time `json:"createdAt"`
	AverageScore float32   `json:"averageScore"`
	Rank         string    `json:"rank"`
	AccountId    int       `json:"accountId"`
}

// Constructor: tạo MarkTable mới với CreatedAt = time.Now()
func NewMarkTable(id int, averageScore float32, rank string, accountId int) *MarkTable {
	return &MarkTable{
		ID:           id,
		CreatedAt:    time.Now(),
		AverageScore: averageScore,
		Rank:         rank,
		AccountId:    accountId,
	}
}

// Getters
func (m *MarkTable) GetID() int               { return m.ID }
func (m *MarkTable) GetCreatedAt() time.Time  { return m.CreatedAt }
func (m *MarkTable) GetAverageScore() float32 { return m.AverageScore }
func (m *MarkTable) GetRank() string          { return m.Rank }
func (m *MarkTable) GetAccountId() int        { return m.AccountId }

// Setters
func (m *MarkTable) SetID(id int)                  { m.ID = id }
func (m *MarkTable) SetCreatedAt(t time.Time)      { m.CreatedAt = t }
func (m *MarkTable) SetAverageScore(score float32) { m.AverageScore = score }
func (m *MarkTable) SetRank(rank string)           { m.Rank = rank }
func (m *MarkTable) SetAccountId(accountId int)    { m.AccountId = accountId }

// String method
func (m *MarkTable) String() string {
	return fmt.Sprintf(
		"MarkTable[ID=%d, CreatedAt=%s, AverageScore=%.2f, Rank=%s, AccountId=%d]",
		m.ID,
		m.CreatedAt.Format("2006-01-02 15:04:05"),
		m.AverageScore,
		m.Rank,
		m.AccountId,
	)
}
