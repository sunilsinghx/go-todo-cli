package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

// todo struct
type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

// todo slice
type Todos []Todo

// ADD
func (todos *Todos) add(title string) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CompletedAt: nil,
		CreatedAt:   time.Now(),
	}

	//append todo to Todos
	*todos = append(*todos, todo)
}

// validate index
func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("invalid index")
		fmt.Println(err.Error())
		return err
	}
	return nil
}

// delete
func (todos *Todos) delete(index int) error {

	t := *todos
	if err := t.validateIndex(index); err != nil {
		return err
	}

	//it will append all todos except current
	*todos = append(t[:index], t[index+1:]...)

	return nil
}

// toggle i.e to make todo as completed/incomplete
func (todos *Todos) toggle(index int) error {

	t := *todos

	if err := t.validateIndex(index); err != nil {
		return err
	}

	isCompleted := t[index].Completed

	//update completion time
	if !isCompleted {
		completionTime := time.Now()
		t[index].CompletedAt = &completionTime
	}
	//if todo is not completed mark it Complete vice verca
	t[index].Completed = !isCompleted

	return nil
}

// edit
func (todos *Todos) edit(index int, title string) error {
	t := *todos

	if err := t.validateIndex(index); err != nil {
		return err
	}
	//update title
	t[index].Title = title

	return nil
}

// Print
func (todos *Todos) print() {
	table := table.New(os.Stdout)
	table.SetRowLines(true)
	table.SetHeaders("#", "Title", "Completed", "Created At", "Completed At")

	for index, t := range *todos {
		completed := "❌"
		completedAt := ""

		//if todos completed assign compleetd as ✅ and
		//make completedAt in readable format
		if t.Completed {
			completed = "✅"

			if t.CompletedAt != nil {
				completedAt = t.CompletedAt.Format(time.RFC1123)
			}
		}

		table.AddRow(strconv.Itoa(index), t.Title, completed, t.CreatedAt.Format(time.RFC1123), completedAt)
	}
	table.Render()
}
