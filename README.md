# Todo CLI in Go

This is a simple command-line TODO list manager built in Go. It allows you to add, delete, edit, toggle, and list your todos. All todos are saved in a JSON file for persistence.

## Features
- **Add** a new todo.
- **Delete** a todo by index.
- **Edit** a todo by index.
- **Toggle** the completion status of a todo.
- **List** all todos.

## Usage

Run the CLI tool with `go run ./ ` and specify the desired command.

```bash
go run ./ [flags]
```

## Add a new Todo
Example:

```bash
go run ./ -add "Buy groceries"
```

## Delete a Todo

```bash
go run ./ -del 1
```

## Edit a Todo


```bash
go run ./ -edit "1:Buy milk"
```


## List all Todos


```bash
go run ./ -list
```

## Toggle Todo completion


```bash
go run ./ -toggle 0
```
