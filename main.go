package main

import (
	"fmt"
	"student_management_be/models"
	"time" // ✅ needed for time.Time
)

func main() {
	tt := models.TimeTable{
		ID:   1,
		Name: "Morning Schedule",
	}

	class := models.Classroom{
		ID:          101,
		Name:        "Physics",
		Total:       45,
		TimeTableID: tt.ID,
	}

	// Correct: Birth is time.Time, not string
	birthDate, _ := time.Parse("2006-01-02", "2025-09-01")

	account := models.Account{
		ID:          1,
		Name:        "Cookiescooker",
		Email:       "no@cookiescooker.click",
		Birth:       birthDate,
		Role:        "student",
		TimeTableID: tt.ID,
	}

	fmt.Println(tt)      // TimeTable[ID=1, Name=Morning Schedule]
	fmt.Println(class)   // Classroom[ID=101, Name=Physics, Total=45, TimeTableID=1]
	fmt.Println(account) // Account[ID=1, Name=Cookiescooker, Email=no@..., Birth=2025-09-01, Role=student, TimeTableID=1]
}
