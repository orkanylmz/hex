package v0

// Generate gRPC client and server auto-generated code
//go:generate protoc --proto_path=. --go_out=. --go_opt=module=hex/internal/grpc/pb/v0 user.proto
//go:generate protoc --proto_path=. --go-grpc_out=. --go-grpc_opt=module=hex/internal/grpc/pb/v0 user.proto
