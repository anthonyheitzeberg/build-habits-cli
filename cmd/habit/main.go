package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

type HabitTracker struct {
	Habits []string
	Grid   map[string][]bool
	Month  time.Month
	Year   int
	Days   int
}

func NewHabitTracker(habits []string, month time.Month, year int) *HabitTracker {
	days := daysInMonth(month, year)
	grid := make(map[string][]bool)
	for _, habit := range habits {
		grid[habit] = make([]bool, days)
	}
	return &HabitTracker{
		Habits: habits,
		Grid:   grid,
		Month:  month,
		Year:   year,
		Days:   days,
	}
}

func daysInMonth(month time.Month, year int) int {
	return time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC).Day()
}

func clearScreen() {
	cmd := exec.Command("clear")
	if os.Getenv("OS") == "Windows_NT" {
		cmd = exec.Command("cmd", "/c", "cls")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func (ht *HabitTracker) DisplayGrid() {
	fmt.Printf("Habit Tracker for %s %d\n\n", ht.Month, ht.Year)
	fmt.Printf("%-15s", "")
	for day := 1; day <= ht.Days; day++ {
		fmt.Printf(" %-2d", day)
	}
	fmt.Println()

	for _, habit := range ht.Habits {
		fmt.Printf("%-15s", habit) // Align habit names
		for day := 1; day <= ht.Days; day++ {
			if ht.Grid[habit][day-1] {
				fmt.Print(" ✓ ")
			} else {
				fmt.Print(" . ")
			}
		}
		fmt.Println()
	}
}

func (ht *HabitTracker) MarkHabit(habit string, day int) {
	if day < 1 || day > ht.Days {
		fmt.Println("Invalid day")
		return
	}
	if _, exists := ht.Grid[habit]; exists {
		ht.Grid[habit][day-1] = true
		fmt.Printf("Marked %s for day %d\n", habit, day)
	} else {
		fmt.Println("Habit does not exist")
	}
}

func main() {
	habits := []string{"Exercise", "Read", "Meditate"}
	month := time.January
	year := 2024

	tracker := NewHabitTracker(habits, month, year)

	for {
		clearScreen()
		tracker.DisplayGrid()
		fmt.Println("\nOptions:")
		fmt.Println("1. Mark habit")
		fmt.Println("2. Quit")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Print("Enter habit: ")
			var habit string
			fmt.Scan(&habit)
			if _, exists := tracker.Grid[habit]; !exists {
				tracker.Habits = append(tracker.Habits, habit)
				tracker.Grid[habit] = make([]bool, tracker.Days)
				fmt.Printf("Added new habit: %s\n", habit)
			}
			fmt.Print("Enter day: ")
			var day int
			fmt.Scan(&day)
			tracker.MarkHabit(habit, day)
		case 2:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option!")
		}
	}
}
