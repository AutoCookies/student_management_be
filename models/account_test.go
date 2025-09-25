package models

import (
	"testing"
	"time"
)

func TestAccount(t *testing.T) {
	// Khởi tạo dữ liệu
	birth := time.Date(1995, 7, 10, 0, 0, 0, 0, time.UTC)
	timetableID := 101
	acc := NewAccount(1, "Alice", "alice@example.com", birth, "Admin", &timetableID)

	// --- Test Getter ---
	if acc.GetName() != "Alice" {
		t.Errorf("Expected Name 'Alice', got '%s'", acc.GetName())
	}

	if acc.GetEmail() != "alice@example.com" {
		t.Errorf("Expected Email 'alice@example.com', got '%s'", acc.GetEmail())
	}

	if acc.GetRole() != "Admin" {
		t.Errorf("Expected Role 'Admin', got '%s'", acc.GetRole())
	}

	if !acc.GetBirth().Equal(birth) {
		t.Errorf("Expected Birth '%v', got '%v'", birth, acc.GetBirth())
	}

	if acc.GetTimeTableID() == nil || *acc.GetTimeTableID() != 101 {
		t.Errorf("Expected TimeTableID 101, got %v", acc.GetTimeTableID())
	}

	if acc.GetCreatedAt().IsZero() {
		t.Error("Expected CreatedAt to be set, got zero value")
	}

	// --- Test Setter ---
	acc.SetName("Alice Nguyen")
	if acc.GetName() != "Alice Nguyen" {
		t.Errorf("Expected Name 'Alice Nguyen', got '%s'", acc.GetName())
	}

	acc.SetEmail("alice.nguyen@example.com")
	if acc.GetEmail() != "alice.nguyen@example.com" {
		t.Errorf("Expected Email 'alice.nguyen@example.com', got '%s'", acc.GetEmail())
	}

	acc.SetRole("SuperAdmin")
	if acc.GetRole() != "SuperAdmin" {
		t.Errorf("Expected Role 'SuperAdmin', got '%s'", acc.GetRole())
	}

	newBirth := time.Date(1996, 8, 15, 0, 0, 0, 0, time.UTC)
	acc.SetBirth(newBirth)
	if !acc.GetBirth().Equal(newBirth) {
		t.Errorf("Expected Birth '%v', got '%v'", newBirth, acc.GetBirth())
	}

	newTimeTableID := 202
	acc.SetTimeTableID(&newTimeTableID)
	if acc.GetTimeTableID() == nil || *acc.GetTimeTableID() != 202 {
		t.Errorf("Expected TimeTableID 202, got %v", acc.GetTimeTableID())
	}

	// --- Test String ---
	str := acc.String()
	if str == "" {
		t.Error("Expected String() to return non-empty string")
	}
}
