## Overview

This project implements a simple gRPC service with a "Greeter" service that provides a `SayHello` RPC method. It demonstrates:

- Protocol Buffers (protobuf) definition
- gRPC server implementation with reflection enabled
- gRPC client implementation
- Code generation from `.proto` files


## Project Structure

```
grpc-basics-go/
├── cmd/
│   ├── client/
│   │   └── main.go          # gRPC client implementation
│   └── server/
│       └── main.go          # gRPC server implementation
├── gen/
│   └── greeter/
│       └── v1/
│           ├── greeter.pb.go        # Generated message types
│           └── greeter_grpc.pb.go   # Generated gRPC service code
├── proto/
│   └── greeter/
│       └── v1/
│           └── greeter.proto        # Protocol Buffer definition
├── go.mod
└── go.sum
```

## Setup

1. **Clone or navigate to the project directory:**
   ```bash
   cd grpc-basics-go
   ```

2. **Install Go dependencies:**
   ```bash
   go mod tidy
   ```

## Generating Code from Protocol Buffers

The `gen/` directory contains generated Go code from the `.proto` files. To regenerate this code after modifying the proto files:

```bash
cd grpc-basics-go

protoc \
  --go_out=. \
  --go_opt=paths=source_relative \
  --go-grpc_out=. \
  --go-grpc_opt=paths=source_relative \
  proto/greeter/v1/greeter.proto
```

This command:
- Generates Go code for message types (`greeter.pb.go`)
- Generates Go code for gRPC services (`greeter_grpc.pb.go`)
- Outputs files to `gen/greeter/v1/` as specified in the `go_package` option

**Note:** The generated files are already present in the repository, so you only need to run this if you modify the `.proto` file.

## Running the Server

Start the gRPC server on port `50051`:

```bash
cd grpc-basics-go
go run cmd/server/main.go
```

You should see:
```
gRPC server listening on port 50051
```

The server:
- Listens on `localhost:50051`
- Implements the `GreeterService` with `SayHello` RPC
- Has gRPC reflection enabled (useful for tools like `grpcurl` or `bloomrpc`)

## Running the Client

In a separate terminal, run the client to make a request to the server:

```bash
cd grpc-basics-go
go run cmd/client/main.go
```

The client will:
- Connect to the server at `localhost:50051`
- Send a `SayHello` request with name "World"
- Print the response: `SayHello response: Hello, World!`

**Note:** Make sure the server is running before starting the client.

## Building Executables

You can build standalone executables:

```bash
# Build server
go build -o bin/server ./cmd/server

# Build client
go build -o bin/client ./cmd/client

# Run the built executables
./bin/server
./bin/client
```

## Protocol Buffer Definition

The service is defined in `proto/greeter/v1/greeter.proto`:

```protobuf
service GreeterService {
    rpc SayHello(SayHelloRequest) returns (SayHelloResponse);
}

message SayHelloRequest {
    string name = 1;
}

message SayHelloResponse {
    string message = 1;
}
```

## Testing with gRPC Tools

Since reflection is enabled, you can use tools like `grpcurl` to interact with the server:

1. **Install grpcurl:**
   ```bash
   go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
   ```

2. **List services:**
   ```bash
   grpcurl -plaintext localhost:50051 list
   ```

3. **Call the SayHello method:**
   ```bash
   grpcurl -plaintext -d '{"name": "Alice"}' \
     localhost:50051 \
     greeter.v1.GreeterService/SayHello
   ```

## Implementation Details

### Server (`cmd/server/main.go`)
- Implements `greeterv1.GreeterServiceServer` interface
- Handles `SayHello` requests and responds with a greeting message
- Uses gRPC reflection for service discovery
- Defaults name to "World" if empty

### Client (`cmd/client/main.go`)
- Creates an insecure connection (for development only)
- Uses a 2-second timeout context
- Sends a request and logs the response

## Dependencies

Key dependencies (managed via `go.mod`):
- `google.golang.org/grpc` - gRPC framework
- `google.golang.org/protobuf` - Protocol Buffers support

## Troubleshooting

**Issue: "could not import google.golang.org/grpc/reflection"**
- Solution: Run `go mod tidy` to download dependencies

**Issue: "protoc: command not found"**
- Solution: Install Protocol Buffers compiler (see Prerequisites)

**Issue: "protoc-gen-go: program not found"**
- Solution: Install the Go plugins (see Prerequisites) and ensure they're in your PATH

**Issue: "connection refused"**
- Solution: Make sure the server is running before starting the client
