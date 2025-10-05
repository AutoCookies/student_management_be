package dto

import (
	"time"
)

type CreateMarkTable struct {
	AverageScore float32 `json:"averageScore"`
	Ranked       string  `json:"ranked"`
	AccountId    int     `json:"accountId"`
}

type MarkTableResponse struct {
	ID           int       `json:"id"`
	CreatedAt    time.Time `json:"createdAt"`
	AverageScore float32   `json:"averageScore"`
	Ranked       string    `json:"ranked"`
	AccountId    int       `json:"accountId"`
}
