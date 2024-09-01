Here's an enhanced README for your Git repository, detailing the structure and usage of your Go-based auto-scaler application:

---

![autoscaler-white](https://github.com/user-attachments/assets/b85a731e-4f0c-4348-afea-c764e2c5eb2e)


# Auto-scaler

A Go application that automatically adjusts the number of replicas of a separate application based on CPU utilization metrics.

## Table of Contents
- [Overview](#overview)
- [Project Structure](#project-structure)
- [Installation](#installation)
- [Usage](#usage)
- [Configuration](#configuration)
- [Building the Application](#building-the-application)
- [Running the Application](#running-the-application)
- [Contributing](#contributing)
- [License](#license)

## Overview

The auto-scaler is a lightweight Go application that dynamically scales the number of replicas of a target application based on its CPU usage. It periodically checks the CPU utilization and adjusts the replicas count to maintain the desired target CPU utilization.

## Project Structure

```
.
├── cmd
│   └── auto-scaler
│       └── main.go               # Entry point of the application
├── internal
│   ├── api
│   │   └── client.go             # API client for interacting with the target application
│   ├── config
│   │   └── config.go             # Configuration parsing and management
│   └── scaler
│       └── scaler.go             # Core auto-scaling logic
├── go.mod                        # Go module file
└── README.md                     # Project documentation
```

### Key Components

- **`cmd/auto-scaler/main.go`**: The main entry point of the application. Handles initialization and starts the auto-scaling process.
- **`internal/config/config.go`**: Contains configuration management logic, including parsing command-line flags.
- **`internal/api/client.go`**: Defines a simple API client to interact with the target application to retrieve metrics and update replicas.
- **`internal/scaler/scaler.go`**: Implements the core logic for checking CPU utilization and adjusting the number of replicas.

## Installation

To install the application, ensure you have Go installed on your system. You can download Go from the official [Go website](https://golang.org/dl/).

1. Clone the repository:

    ```sh
    git clone https://github.com/Thrinadh-Kumpatla/auto-scaler.git
    cd auto-scaler
    ```

2. Install dependencies:

    The `go.mod` file will handle dependencies. Run the following command to install them:

    ```sh
    go mod tidy
    ```

## Usage

To run the auto-scaler, execute the following command:

```sh
go run cmd/auto-scaler/main.go [flags]
```

### Flags

- `--port`: Port of the application (default: 8123)
- `--target-cpu`: Target CPU utilization (0.0-1.0) (default: 0.80)
- `--check-interval`: Interval between checks (default: 10s)

Example:

```sh
go run cmd/auto-scaler/main.go --port 8123 --target-cpu 0.75 --check-interval 15s
```

This example starts the auto-scaler, targeting a CPU utilization of 75% and checking every 15 seconds.

## Configuration

The application can be configured using command-line flags:

- **Port (`--port`)**: Specifies the port of the target application.
- **Target CPU Utilization (`--target-cpu`)**: Sets the desired CPU utilization threshold. The auto-scaler will adjust the replicas to maintain this utilization.
- **Check Interval (`--check-interval`)**: Defines how often the auto-scaler checks the CPU usage.

## Building the Application

To build the application, use the Go build command:

```sh
go build -o auto-scaler cmd/auto-scaler/main.go
```

This command creates an executable named `auto-scaler` in the current directory.

## Running the Application

After building, you can run the application with the desired configuration:

```sh
./auto-scaler --port 8123 --target-cpu 0.75 --check-interval 15s
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---
