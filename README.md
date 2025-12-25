# go-grpc

A repository containing multiple gRPC learning projects and examples implemented in Go.

## Overview

This repository serves as a collection of gRPC projects demonstrating various concepts, patterns, and best practices for building gRPC services in Go. Each project is self-contained and focuses on different aspects of gRPC development.

## Prerequisites

Before working with any project in this repository, ensure you have the following installed:

### 1. Go

Install Go (version 1.25.5 or later recommended):

**macOS:**
```bash
brew install go
```

**Linux (Ubuntu/Debian):**
```bash
sudo apt-get update
sudo apt-get install -y golang-go
```

**Or download from:** https://go.dev/dl/

**Verify installation:**
```bash
go version
```

**Set up Go workspace (if needed):**
```bash
# Add to your ~/.zshrc or ~/.bashrc
export GOPATH=$HOME/go
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
```

### 2. Protocol Buffers Compiler (protoc)

The Protocol Buffers compiler is required to generate Go code from `.proto` files.

**macOS:**
```bash
brew install protobuf
```

**Linux (Ubuntu/Debian):**
```bash
sudo apt-get install -y protobuf-compiler
```

**Or download from:** https://grpc.io/docs/protoc-installation/

**Verify installation:**
```bash
protoc --version
```

### 3. Go Plugins for Protocol Buffers

Install the required Go plugins for generating gRPC code:

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

**Ensure plugins are in your PATH:**
```bash
# Add to your ~/.zshrc or ~/.bashrc
export PATH="$PATH:$(go env GOPATH)/bin"

# Or verify the path
echo $(go env GOPATH)/bin
```

**Verify plugins are installed:**
```bash
which protoc-gen-go
which protoc-gen-go-grpc
```

## Projects

### [grpc-basics-go](./grpc-basics-go/)

A basic gRPC project demonstrating fundamental concepts:
- Simple gRPC server and client implementation
- Protocol Buffers definition and code generation
- gRPC reflection for service discovery

See the [project README](./grpc-basics-go/README.md) for detailed setup and usage instructions.

## Repository Structure

```
go-grpc/
├── grpc-basics-go/          # Basic gRPC server and client
│   ├── cmd/                 # Application entry points
│   ├── gen/                 # Generated code from .proto files
│   ├── proto/               # Protocol Buffer definitions
│   └── README.md            # Project-specific documentation
├── LICENSE
└── README.md                # This file
```

## Common Workflows

### Generating Code from Protocol Buffers

Most projects will require generating Go code from `.proto` files. The general pattern is:

```bash
cd <project-directory>

protoc \
  --go_out=. \
  --go_opt=paths=source_relative \
  --go-grpc_out=. \
  --go-grpc_opt=paths=source_relative \
  proto/**/*.proto
```

### Running Projects

Each project has its own `go.mod` file. Navigate to the project directory and:

```bash
# Install dependencies
go mod tidy

# Run the project (see project-specific README)
go run cmd/server/main.go
```

## Development Tools

### Recommended Tools

- **grpcurl** - Command-line tool for interacting with gRPC servers
  ```bash
  go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
  ```

- **BloomRPC** - GUI client for gRPC services
  - Download from: https://github.com/bloomrpc/bloomrpc

- **buf** - Protocol Buffer tooling
  ```bash
  brew install bufbuild/buf/buf
  ```

## Troubleshooting

### Common Issues

**Issue: "protoc: command not found"**
- Solution: Install Protocol Buffers compiler (see Prerequisites section)

**Issue: "protoc-gen-go: program not found"**
- Solution: Install Go plugins and ensure `$GOPATH/bin` is in your PATH

**Issue: "could not import google.golang.org/grpc"**
- Solution: Run `go mod tidy` in the project directory

**Issue: "go: command not found"**
- Solution: Install Go and ensure it's in your PATH

## Contributing

When adding new projects to this repository:

1. Create a new directory for your project
2. Include a `README.md` with project-specific documentation
3. Use a separate `go.mod` file for each project
4. Follow the existing project structure conventions
5. Update this README to include your project in the Projects section

## Resources

- [gRPC Official Documentation](https://grpc.io/docs/)
- [Protocol Buffers Guide](https://protobuf.dev/getting-started/gotutorial/)
- [Go gRPC Tutorial](https://grpc.io/docs/languages/go/basics/)
- [Protocol Buffers Language Guide](https://protobuf.dev/programming-guides/proto3/)

## License

See [LICENSE](./LICENSE) file for details.
