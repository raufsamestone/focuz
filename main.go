package main

import (
	"fmt"
	"time"

	"github.com/gen2brain/beeep"
)

func main() {
	fmt.Println("Welcome to Focuz! Your personal work-break reminder")

	fmt.Print("Enter work duration (in minutes or seconds): ")
	var workDuration time.Duration
	fmt.Scanln(&workDuration)

	fmt.Print("Enter break duration (in minutes or seconds): ")
	var breakDuration time.Duration
	fmt.Scanln(&breakDuration)

	fmt.Println("Press 'i' to interrupt the current work/break session")
	fmt.Println("Press 'd' to mark the current work session as done and exit")
	fmt.Println("Press 'q' to quit the program")

	workSessions := 0
	totalWorkTime := time.Duration(0)
	totalBreakTime := time.Duration(0)
	isWorkTime := true
	interrupted := false

	for {
		if interrupted {
			interrupted = false
			isWorkTime = !isWorkTime
		}

		var currentSessionDuration time.Duration
		var sessionName string
		if isWorkTime {
			currentSessionDuration = workDuration
			sessionName = "work"
		} else {
			currentSessionDuration = breakDuration
			sessionName = "break"
		}

		fmt.Printf("\n%s time started for %v\n", sessionName, currentSessionDuration)
		startTime := time.Now()

	loop:
		for {
			select {
			case <-time.After(currentSessionDuration):
				if isWorkTime {
					workSessions++
					totalWorkTime += currentSessionDuration
				} else {
					totalBreakTime += currentSessionDuration
				}
				break loop
			default:
				if interrupted {
					break loop
				}
			}
		}

		if interrupted {
			fmt.Printf("%s session interrupted.\n", sessionName)
			interrupted = false
			isWorkTime = !isWorkTime
			continue
		}

		endTime := time.Now()
		duration := endTime.Sub(startTime)

		if isWorkTime {
			totalWorkTime += duration
		} else {
			totalBreakTime += duration
		}

		fmt.Printf("%s time's up!\n", sessionName)

		err := beeep.Notify("Focuz", fmt.Sprintf("%s time's up!", sessionName), "assets/warning.png")
		if err != nil {
			panic(err)
		}

		isWorkTime = !isWorkTime

		var input string
		fmt.Scanln(&input)
		switch input {
		case "i":
			interrupted = true
		case "d":
			if isWorkTime {
				workSessions++
				totalWorkTime += duration
			} else {
				totalBreakTime += duration
			}
			fmt.Printf("You have completed %d work sessions with an average work time of %v and an average break time of %v.\n", workSessions, totalWorkTime/time.Duration(workSessions), totalBreakTime/time.Duration(workSessions))
			return
		case "q":
			fmt.Println("Goodbye!")
			return
		}
	}
}
