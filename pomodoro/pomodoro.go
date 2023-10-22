package pomodoro

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func RunPomodoro(workDuration, shortBreakDuration, longBreakDuration int, taskName string) {
	workIntervals := 0

	clearScreen()
	for {
		fmt.Printf("Pomodoro Timer - Working on: %s\n", taskName)
		startTimer(workDuration)
		workIntervals++

		if workIntervals%4 == 0 {
			clearScreen()
			fmt.Println("Long Break")
			startTimer(longBreakDuration)
			clearScreen()
		} else {
			clearScreen()
			fmt.Println("Short Break")
			startTimer(shortBreakDuration)
			clearScreen()
		}
	}
}

func startTimer(duration int) {
	totalSeconds := duration * 60
	seconds := totalSeconds

	for seconds > 0 {
		printProgressBar(totalSeconds-seconds, totalSeconds)
		time.Sleep(1 * time.Second)
		seconds--
	}
}

func clearScreen() {
	cmd := exec.Command("cmd", "/c", "cls") // use clear instead of cls for unix like systems.
	cmd.Stdout = os.Stdout
	_ = cmd.Run()
}

func printProgressBar(current, total int) {
	barLength := 40
	seconds := total - current
	minutes := seconds / 60
	bar := strings.Repeat("=", int(current*barLength/total))
	spaces := strings.Repeat(" ", barLength-len(bar))
	fmt.Printf("[%s%s] %02d:%02d\r", bar, spaces, minutes, seconds%60)
}
