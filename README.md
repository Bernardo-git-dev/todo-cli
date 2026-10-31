# Task Tracker CLI

A lightweight command-line interface (CLI) application built in Go to track and manage your daily tasks. 

This project is an implementation of the [roadmap.sh Task Tracker project idea](https://roadmap.sh/projects/task-tracker).

## Features

The application operates entirely through positional command-line arguments and handles the following operations:

- **Add Tasks**: Create a new task with a unique, auto-incrementing ID.
- **Update Tasks**: Modify the title of an existing task by specifying its ID.
- **Delete Tasks**: Remove a task permanently from the tracking file.
- **Mark Status**: Transition task states natively using dedicated commands:
  - `mark-todo`
  - `mark-in-progress`
  - `mark-done`
- **List & Filter**: View all tasks or narrow down the output by specific statuses (`todo`, `in-progress`, `done`).
- **Automated Data Persistence**: Saves all records inside a local `data.json` file. The file is automatically initialized on the first run if it does not exist.

## Task Data Structure

Each task stored inside `data.json` contains the following structured fields:

```json
{
    "id": "1",
    "title": "Study lexical scope in Go",
    "status": "todo",
    "created_at": "2026-10-31T15:45:00Z",
    "updated_at": "2026-10-31T15:45:00Z"
}
```

## Usage

### Prerequisites
Make sure you have Go installed on your machine.

### Installation
Clone the repository or save the source files locally, then compile the application:
```bash
go build -o task-cli main.go
```

### Commands Examples

**Adding a new task**
```bash
./task-cli add "Buy groceries"
# Output: Task "Buy groceries" added successfully!
```

**Updating a task title**
```bash
./task-cli update 1 "Buy groceries and cook dinner"
```

**Deleting a task**
```bash
./task-cli delete 1
```

**Marking a task status**
```bash
./task-cli mark-in-progress 1
./task-cli mark-done 1
```

**Listing tasks**
```bash
# List all tasks
./task-cli list

# Filter tasks by status
./task-cli list todo
./task-cli list in-progress
./task-cli list done
```

## Architecture & Design Constraints Followed
- **Zero Third-Party Dependencies**: Built entirely using the Go standard library (`os`, `fmt`, `encoding/json`, `time`, `strconv`).
- **Atomic File Creation**: Uses `os.O_CREATE | os.O_EXCL` flags for a secure system-level initialization that prevents race conditions and data corruption.
- **Robust Error Handling**: Defensively checks user inputs (`os.Args`) before index allocation to prevent runtime array panics.
