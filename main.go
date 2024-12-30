package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/go-toast/toast"
)

func main() {
	const workDuration = 25 * time.Minute
	const shortBreak = 5 * time.Minute
	const longBreak = 15 * time.Minute
	const cycles = 4

	for i := 1; i <= cycles; i++ {
		fmt.Printf("Cycle %02d: Word for %02d minutes\n", i, workDuration/time.Minute)
		showNotification("Pomodoro Clock", fmt.Sprintf("Cycle %02d: Work for %02d minutes", i, workDuration/time.Minute))
		startTimer(workDuration)

		if i < cycles {
			shortBreakToString := fmt.Sprint("take a %02d minute break\n", shortBreak/time.Minute)
			showNotification("Pomodoro Clock", shortBreakToString)
			fmt.Println(shortBreakToString)
			startTimer(shortBreak)
		} else {
			longBreakToString := fmt.Sprint("take a %02d minute break\n", longBreak/time.Minute)
			showNotification("Pomodoro Clock", longBreakToString)
			fmt.Println(longBreakToString)
			startTimer(longBreak)
		}
	}

}

func startTimer(duration time.Duration) {
	for duration > 0 {
		fmt.Printf("\rTime left: %02d:%02d", int(duration.Minutes()), int(duration.Seconds())%60)
		time.Sleep(1 * time.Second)
		duration -= time.Second
	}
	fmt.Println("\nTime's up!")
}

func showNotification(title, message string) {

	baseDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current directory: %v", err)
		return
	}

	imgPath := filepath.Join(baseDir, "assets", "clock.png")

	notification := toast.Notification{
		AppID:   "Pomodoro CLock",
		Title:   title,
		Message: message,
		Icon:    imgPath,
	}

	err = notification.Push()

	if err != nil {
		fmt.Printf("Error showcasing notification: %v", err)
	}

}
