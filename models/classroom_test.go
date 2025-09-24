package models

import (
	"testing"
	"time"
)

// Test Classroom
func TestClassroom(t *testing.T) {
	// Constructor
	classroom := NewClassroom(1, "Class A", 30, 101)

	// --- Test Getter ---
	if classroom.GetName() != "Class A" {
		t.Errorf("Expected Name 'Class A', got '%s'", classroom.GetName())
	}

	if classroom.GetTotal() != 30 {
		t.Errorf("Expected Total 30, got %d", classroom.GetTotal())
	}

	if classroom.GetTimeTableID() != 101 {
		t.Errorf("Expected TimeTableID 101, got %d", classroom.GetTimeTableID())
	}

	if classroom.GetCreatedAt().IsZero() {
		t.Error("Expected CreatedAt to be set, got zero value")
	}

	// --- Test Setter ---
	classroom.SetName("Class B")
	if classroom.GetName() != "Class B" {
		t.Errorf("Expected Name 'Class B', got '%s'", classroom.GetName())
	}

	classroom.SetTotal(35)
	if classroom.GetTotal() != 35 {
		t.Errorf("Expected Total 35, got %d", classroom.GetTotal())
	}

	classroom.SetTimeTableID(202)
	if classroom.GetTimeTableID() != 202 {
		t.Errorf("Expected TimeTableID 202, got %d", classroom.GetTimeTableID())
	}

	newCreated := time.Date(2025, 9, 22, 12, 0, 0, 0, time.UTC)
	classroom.SetCreatedAt(newCreated)
	if !classroom.GetCreatedAt().Equal(newCreated) {
		t.Errorf("Expected CreatedAt %v, got %v", newCreated, classroom.GetCreatedAt())
	}

	// --- Test String ---
	str := classroom.String()
	if str == "" {
		t.Error("Expected String() to return non-empty string")
	}
}
