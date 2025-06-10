# Dynamic Worker Pool

Implement a primitive worker-pool with the ability to dynamically add and remove workers. Input data (strings) are sent to a channel, workers process them (for example, display the worker number and the data itself). A task for basic knowledge of channels and goroutines.

## Features

- Dynamically add and remove workers while the program is running
- Workers process string jobs from a shared channel
- Safe for concurrent use
- Unit tests with race detector support

## Project Structure

```
.
├── cmd/
│   └── main.go           # Entry point
├── internal/
│   ├── dispatcher/
│   │   ├── dispatcher.go # Dispatcher logic
│   │   └── dispatcher_test.go # Unit tests
│   └── worker/
│       └── worker.go     # Worker logic
├── README.md
└── .gitignore
```

## Build

```bash
go build -o main ./cmd/main.go
```

## Run

```bash
./main
```
or
```bash
go run ./cmd/main.go
```

## Test

Run all unit tests:

```bash
go test ./internal/dispatcher
```

Run tests with the Go race detector:

```bash
go test -race ./internal/dispatcher
```

## Example Output

```
Worker 1 processing: Job 1 payload
Worker 2 processing: Job 2 payload
Adding a worker...
Worker 3 processing: Job 6 payload
Removing a worker...
Worker 1 processing: Job 9 payload
...
Worker 2 stopping
Worker 3 stopping
```

## How It Works

- The dispatcher maintains a shared job queue and a dynamic list of workers.
- Each worker listens for jobs on the shared queue and processes them.
- You can add or remove workers at runtime; removed workers stop gracefully.

## Contributing

Feel free to open issues or submit pull requests for improvements or bug fixes.

---

**Author:**  
Alexander Alexandrov  
batareyka.work@gmail.com
