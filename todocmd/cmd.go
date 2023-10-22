package todocmd

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/thestuti/devup/todo"
)

func AddTask(todos *todo.Todos, args []string) {
	addCmd := flag.NewFlagSet("tadd", flag.ExitOnError)
	addTask := addCmd.String("task", "", "The content of new todo item")
	addCat := addCmd.String("cat", "Uncategorized", "The category of the todo item")

	addCmd.Parse(args)

	if len(*addTask) == 0 {
		fmt.Println("Error: the --task flag is required for the 'todo add' subcommand.")
		os.Exit(1)
	}

	todos.Add(*addTask, *addCat)
	err := todos.Store(GetJsonFile())
	if err != nil {
		log.Fatal(err)
	}

	todos.Print(2, "")
	fmt.Println("Todo item added successfully.")
}

func DeleteTask(todos *todo.Todos, args []string) {
	deleteCmd := flag.NewFlagSet("tdelete", flag.ExitOnError)

	deleteID := deleteCmd.Int("id", 0, "The id of todo to be deleted")

	deleteCmd.Parse(args)

	err := todos.Delete(*deleteID)
	if err != nil {
		log.Fatal(err)
	}

	err = todos.Store(GetJsonFile())
	if err != nil {
		log.Fatal(err)
	}

	todos.Print(2, "")
	fmt.Println("Todo item deleted successfully.")
}

func Init() {
	ok := GetUserApproval()
	if !ok {
		fmt.Print("You've declined to create the JSON file. You can always run \"tinit\" subcommand again if you change your mind.")
		os.Exit(0)
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	filepath := filepath.Join(homeDir, ".todos.json")
	_, err = os.Stat(filepath)
	if err != nil {
		if os.IsNotExist(err) {
			file, err := os.Create(filepath)
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()
			fmt.Println("Succefully create a \".todos.json\" file in your home directory.")
		} else {
			log.Fatal("Unknown error occurred.")
		}
	} else {
		fmt.Print(".todos.json file exists in your home directory already.")
	}
}

func ListTasks(todos *todo.Todos, args []string) {
	listCmd := flag.NewFlagSet("tlist", flag.ExitOnError)
	listDone := listCmd.Int("done", 2, "The status of todo to be printed")
	listCat := listCmd.String("cat", "", "The category of tasks to be listed")

	listCmd.Parse(args)
	todos.Print(*listDone, *listCat)
}

func UpdateTask(todos *todo.Todos, args []string) {
	updateCmd := flag.NewFlagSet("tupdate", flag.ExitOnError)
	updateID := updateCmd.Int("id", 0, "The id of todo to be updated")
	updateCat := updateCmd.String("cat", "", "The to-be-updated category of todo")
	updateTask := updateCmd.String("task", "", "To to-be-updated content of todo")
	updateDone := updateCmd.Int("done", 2, "The to-be-updated status of todo")

	updateCmd.Parse(args)

	if *updateID == 0 {
		fmt.Println("Error: the --id flag is required for the 'update' subcommand.")
		os.Exit(1)
	}
	err := todos.Update(*updateID, *updateTask, *updateCat, *updateDone)
	if err != nil {
		log.Fatal(err)
	}

	err = todos.Store(GetJsonFile())
	if err != nil {
		log.Fatal(err)
	}

	todos.Print(2, "")
	fmt.Println("Todo item updated successfully.")
}

func GetJsonFile() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	return filepath.Join(homeDir, ".todos.json")
}

func GetUserApproval() bool {
	confirmMessage := "Need to create an empty \".todos.json\" file in your home directory to store your todo items, continue? (y/n): "

	r := bufio.NewReader(os.Stdin)
	var s string

	fmt.Print(confirmMessage)
	s, _ = r.ReadString('\n')
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)

	for {
		if s == "y" || s == "yes" {
			return true
		}
		if s == "n" || s == "no" {
			return false
		}
	}
}

func RemindInit(todos *todo.Todos) {
	_, err := os.Stat(GetJsonFile())
	if err != nil {
		fmt.Println("Please run \"tinit\" subcommand to create an JSON file to store your todo items.")
		os.Exit(1)
	} else {
		if err := todos.Load(GetJsonFile()); err != nil {
			log.Fatal(err)
		}
	}
}
