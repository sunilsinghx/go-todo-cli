package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// CmdFlags struct holds command-line flags for Todo operations
type CmdFlags struct {
	Add    string // Title of the new Todo to add
	Del    int    // Index of the Todo to delete
	Edit   string // Edit a Todo by index & specify a new title
	Toggle int    // Index of the Todo to toggle its completed state
	List   bool   // Flag to list all Todos
}

// NewCmdFlags initializes the command-line flags and returns a CmdFlags instance
func NewCmdFlags() *CmdFlags {
	// Initialize CmdFlags
	cf := CmdFlags{}

	// Define flags for various Todo operations
	flag.StringVar(&cf.Add, "add", "", "Add a new Todo specify title")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a todo by index & specify a new title. id:new_title")
	flag.IntVar(&cf.Del, "del", -1, "Specify todo by index to delete")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Specify todo by index to toggle complete true/false")
	flag.BoolVar(&cf.List, "list", false, "List all todos")

	// Parse command-line flags
	flag.Parse()

	return &cf
}

// Execute processes the command based on the flags set
func (cf *CmdFlags) Execute(todos *Todos) {
	switch {
	// Case to list all Todos
	case cf.List:
		todos.print()

	// Case to add a new Todo
	case cf.Add != "":
		todos.add(cf.Add)

	// Case to edit an existing Todo
	case cf.Edit != "":
		// Split the edit flag into index and new title
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			// Handle invalid edit format
			fmt.Println("Error: Invalid format for edit. Please use index:new_title")
			os.Exit(1)
		}

		// Convert the index string to an integer
		index, err := strconv.Atoi(parts[0])
		if err != nil {
			// Handle invalid index
			fmt.Println("Error: Invalid index for edit")
			os.Exit(1)
		}

		// Edit the Todo at the specified index
		todos.edit(index, parts[1])

	// Case to toggle the completion state of a Todo
	case cf.Toggle != -1:
		todos.toggle(cf.Toggle)

	// Case to delete a Todo by index
	case cf.Del != -1:
		todos.delete(cf.Del)

	// Default case for invalid commands
	default:
		fmt.Println("Invalid Command")
	}
}
