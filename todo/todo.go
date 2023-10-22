package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/alexeyco/simpletable"
)

type item struct {
	ID          int
	Task        string
	Category    string
	Done        bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []item

var nextID int

func (t *Todos) Add(task string, cat string) {
	todo := item{
		ID:          nextID,
		Task:        task,
		Category:    cat,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}

	nextID++

	*t = append(*t, todo)
}

func (t *Todos) Update(id int, task string, cat string, done int) error {
	ls := *t

	index := t.getIndexByID(id)
	if index == -1 {
		return errors.New("invalid ID")
	}

	if len(task) != 0 {
		ls[index].Task = task
	}

	if len(cat) != 0 {
		ls[index].Category = cat
	}

	if done == 0 {
		ls[index].Done = false
		ls[index].CompletedAt = nil
	} else if done == 1 {
		ls[index].Done = true
		completedAt := time.Now()
		ls[index].CompletedAt = &completedAt
	}

	return nil
}

func (t *Todos) Delete(id int) error {
	ls := *t
	index := t.getIndexByID(id)
	if index == -1 {
		return errors.New("invalid ID")
	}

	*t = append(ls[:index], ls[index+1:]...)

	return nil
}

func (t *Todos) Load(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	if len(data) == 0 {
		return err
	}

	err = json.Unmarshal(data, t)
	if err != nil {
		return err
	}

	if len(*t) > 0 {
		maxID := (*t)[0].ID
		for _, todo := range *t {
			if todo.ID > maxID {
				maxID = todo.ID
			}
		}
		nextID = maxID + 1
	}

	return nil
}

func (t *Todos) Store(filename string) error {
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}

func (t *Todos) Print(status int, cat string) {
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Category"},
			{Align: simpletable.AlignCenter, Text: "Task"},
			{Align: simpletable.AlignCenter, Text: "Done?"},
			{Align: simpletable.AlignRight, Text: "CreatedAt"},
			{Align: simpletable.AlignRight, Text: "CompletedAt"},
		},
	}

	var cells [][]*simpletable.Cell

	requestedTodos := []item{}

	for _, todo := range *t {
		if status == 1 {
			if todo.Done {
				requestedTodos = append(requestedTodos, todo)
			}
		}

		if status == 0 {
			if !todo.Done {
				requestedTodos = append(requestedTodos, todo)
			}
		}

		if status != 1 && status != 0 {
			requestedTodos = append(requestedTodos, todo)
		}
	}

	requestedCatTodos := []item{}

	for _, todo := range requestedTodos {
		if strings.EqualFold(todo.Category, cat) || cat == "" {
			requestedCatTodos = append(requestedCatTodos, todo)
		}
	}

	for _, item := range requestedCatTodos {
		task := item.Task
		done := "No"
		completedAt := ""

		if item.Done {
			task = fmt.Sprintf(item.Task)
			done = "\u2705"
		}

		if item.CompletedAt != nil {
			completedAt = item.CompletedAt.Format("2006-01-02")
		}

		cells = append(cells, []*simpletable.Cell{
			{Text: fmt.Sprintf("%d", item.ID)},
			{Text: item.Category},
			{Text: task},
			{Text: done},
			{Text: item.CreatedAt.Format("2006-01-02")},
			{Text: completedAt},
		})
	}

	table.Body = &simpletable.Body{Cells: cells}

	table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
		{Align: simpletable.AlignLeft, Text: ""},
		{Align: simpletable.AlignLeft, Span: 5, Text: fmt.Sprintf("You have %d pending todos", t.CountPending())},
	}}

	table.SetStyle(simpletable.StyleUnicode)

	table.Println()
}

func (t *Todos) CountPending() int {
	total := 0
	for _, item := range *t {
		if !item.Done {
			total++
		}
	}

	return total
}

func (t *Todos) getIndexByID(id int) int {
	index := -1
	for i, todo := range *t {
		if todo.ID == id {
			index = i
			break
		}
	}
	return index
}
