package models

import (
	"fmt"
)

type MarkDetails struct {
	ID          int     `json:"id"`
	TestId      int     `json:"testId"`
	Mark        float32 `json:"mark"`
	MarkTableId int     `json:"markTableId"`
}

// Constructor
func NewMarkDetails(id int, testId int, mark float32, markTableId int) *MarkDetails {
	return &MarkDetails{
		ID:          id,
		TestId:      testId,
		Mark:        mark,
		MarkTableId: markTableId,
	}
}

// Getters
func (m *MarkDetails) GetID() int          { return m.ID }
func (m *MarkDetails) GetTestId() int      { return m.TestId }
func (m *MarkDetails) GetMark() float32    { return m.Mark }
func (m *MarkDetails) GetMarkTableId() int { return m.MarkTableId }

// Setters
func (m *MarkDetails) SetID(id int)                   { m.ID = id }
func (m *MarkDetails) SetTestId(testId int)           { m.TestId = testId }
func (m *MarkDetails) SetMark(mark float32)           { m.Mark = mark }
func (m *MarkDetails) SetMarkTableId(markTableId int) { m.MarkTableId = markTableId }

// String method
func (m *MarkDetails) String() string {
	return fmt.Sprintf(
		"MarkDetails[ID=%d, TestId=%d, Mark=%.2f, MarkTableId=%d]",
		m.ID,
		m.TestId,
		m.Mark,
		m.MarkTableId,
	)
}
