# shared-lib

A Go shared library demonstrating Go workspace functionality with 2 modules.

## Overview

This repository contains a Go workspace with two modules:

1. **utils** - A shared library module containing common utilities (math and string operations)
2. **consumer** - A consumer module that imports and uses the utils module

## Structure

```
├── go.work                 # Go workspace file
├── utils/                  # Shared utilities module
│   ├── go.mod
│   ├── math.go            # Math utility functions
│   ├── math_test.go       # Tests for math functions
│   ├── string.go          # String utility functions
│   └── string_test.go     # Tests for string functions
└── consumer/               # Consumer module
    ├── go.mod
    └── main.go            # Example usage of utils module
```

## Usage

### Running the Consumer Application

```bash
cd consumer
go run main.go
```

This will demonstrate the usage of the shared utilities:

```
Add(10, 20) = 30
Multiply(10, 20) = 200
Max(10, 20) = 20
Capitalize('hello world') = 'Hello world'
Reverse('hello world') = 'dlrow olleh'
```

### Running Tests

To test the utilities module:

```bash
cd utils
go test -v
```

### Building Modules

Build the utils library:
```bash
cd utils
go build
```

Build the consumer application:
```bash
cd consumer
go build
```

### Workspace Commands

Synchronize workspace dependencies:
```bash
go work sync
```

Add a new module to the workspace:
```bash
go work use ./path/to/module
```

## Go Workspace Benefits

This setup demonstrates how Go workspaces allow:

1. **Local module development** - The consumer module can import the utils module directly from the local filesystem
2. **Cross-module refactoring** - Changes in utils are immediately available to consumer without publishing
3. **Unified dependency management** - The workspace manages dependencies across all modules
4. **Easy testing** - Test changes across multiple related modules simultaneously
