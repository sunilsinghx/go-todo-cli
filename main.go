package main

func main() {
	// Initialize an empty Todos slice to hold our Todo items
	todos := Todos{}

	// Create a new Storage instance to handle loading and saving of Todos from/to a file
	// The "todos.json" file is where the Todo list will be stored
	storage := NewStorage[Todos]("todos.json")

	// Load the Todos from the storage (e.g., "todos.json" file) into the 'todos' slice
	// If the file exists, this will populate the 'todos' slice with existing data
	storage.Load(&todos)

	// Initialize command-line flags for handling user input
	// The user will interact with the program by passing flags (e.g., -add, -edit)
	CmdFlags := NewCmdFlags()

	// Execute the command based on the parsed flags
	// This modifies the 'todos' slice by adding, deleting, editing, or toggling Todo items
	CmdFlags.Execute(&todos)

	// After executing the command, save the updated Todos back to the storage (e.g., "todos.json")
	storage.Save(todos)
}
