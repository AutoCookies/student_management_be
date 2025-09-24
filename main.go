package main

import (
	"fmt"
	"student_management_be/models"
	"time"
)

func main() {
	// ---------- Test Account ----------
	fmt.Println("===== Account =====")
	birth := time.Date(1995, 7, 10, 0, 0, 0, 0, time.UTC)
	acc := models.NewAccount(1, "Alice", "alice@example.com", birth, "Admin", 101)
	fmt.Println(acc.String())

	acc.SetName("Alice Nguyen")
	acc.SetEmail("alice.nguyen@example.com")
	fmt.Println("After update:", acc.String())

	// ---------- Test TimeTable ----------
	fmt.Println("\n===== TimeTable =====")
	tt := models.NewTimeTable(1, "Morning Schedule")
	fmt.Println(tt.String())

	tt.SetName("Evening Schedule")
	fmt.Println("After update:", tt.String())

	// ---------- Test MarkTable ----------
	fmt.Println("\n===== MarkTable =====")
	markTable := models.NewMarkTable(1, 8.75, "A", 1)
	fmt.Println(markTable.String())

	markTable.SetAverageScore(9.2)
	markTable.SetRank("A+")
	fmt.Println("After update:", markTable.String())

	// ---------- Test Subject ----------
	fmt.Println("\n===== Subject =====")
	subject := models.NewSubject(1, "Math", 1)
	fmt.Println(subject.String())

	subject.SetName("Advanced Math")
	subject.SetSemester(2)
	fmt.Println("After update:", subject.String())

	// ---------- Test Test ----------
	fmt.Println("\n===== Test =====")
	test := models.NewTest(1, 1)
	fmt.Println(test.String())

	test.SetAccountId(2)
	fmt.Println("After update:", test.String())

	// ---------- Test MarkDetails ----------
	fmt.Println("\n===== MarkDetails =====")
	markDetails := models.NewMarkDetails(1, 1, 9.5, 1)
	fmt.Println(markDetails.String())

	markDetails.SetMark(9.8)
	markDetails.SetTestId(2)
	fmt.Println("After update:", markDetails.String())
}
