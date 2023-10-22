package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/soffyzen/devup/pomodoro"
	"github.com/soffyzen/devup/todo"
	"github.com/soffyzen/devup/todocmd"
)

func Help() {
	fmt.Println("Welcome to devup CLI!")
	fmt.Println("Usage: devup <app> <command> [arguments]")
	fmt.Println()
	fmt.Println("Available commands:")
	fmt.Println()
	fmt.Println("Todo App:")
	fmt.Println("  tinit                                            Create an empty JSON file to store tasks")
	fmt.Println("  tadd --task <task> --cat <cat>                   Add a new task")
	fmt.Println("  tlist <done>                                     List all tasks")
	fmt.Println("  tupdate --task <task> --cat <cat> --id <id>      Update an existing task")
	fmt.Println("  tdelete --id <id>                                Delete an existing task")
	fmt.Println("  tupdate --id <id> --done 1                       Mark as done")
	fmt.Println("")
	fmt.Println("Pomodoro:")
	fmt.Println("  pomodoro -w=<work_duration> -s=<short_break> -l=<long_break> -n=<task_name>")
	fmt.Println("    -w: Working session duration in minutes (default: 25)")
	fmt.Println("    -s: Short break duration in minutes (default: 5)")
	fmt.Println("    -l: Long break duration in minutes (default: 15)")
	fmt.Println("    -n: Task name (default: 'Task')")
}

func main() {
	if len(os.Args) < 2 {
		Help()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "help":
		Help()
	case "tinit":
		todocmd.Init()
	case "tadd", "tlist", "tdelete", "tupdate":
		todoRun()
	case "pomodoro":
		pomodoroRun()
	default:
		fmt.Println("Invalid command. Please use 'help' command to see available commands.")
		os.Exit(1)
	}
}

func todoRun() {
	todos := &todo.Todos{}

	switch os.Args[1] {
	case "tadd":
		todocmd.RemindInit(todos)
		todocmd.AddTask(todos, os.Args[2:])
	case "tlist":
		todocmd.RemindInit(todos)
		todocmd.ListTasks(todos, os.Args[2:])
	case "tdelete":
		todocmd.RemindInit(todos)
		todocmd.DeleteTask(todos, os.Args[2:])
	case "tupdate":
		todocmd.RemindInit(todos)
		todocmd.UpdateTask(todos, os.Args[2:])
	default:
		fmt.Println("Invalid command for the Todo App.")
		os.Exit(1)
	}
}

func pomodoroRun() {
	fs := flag.NewFlagSet("pomodoro", flag.ExitOnError)
	workDuration := fs.Int("w", 25, "Working session duration (in minutes)")
	shortBreakDuration := fs.Int("s", 5, "Short break duration (in minutes)")
	longBreakDuration := fs.Int("l", 15, "Long break duration (in minutes)")
	taskName := fs.String("n", "Task", "Task name")

	fs.Parse(os.Args[2:])

	pomodoro.RunPomodoro(*workDuration, *shortBreakDuration, *longBreakDuration, *taskName)
}
