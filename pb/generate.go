package pb

//go:generate protoc -I . --go_out=. --go_opt=module=github.com/bbengfort/oneof/pb --go-grpc_out=. --go-grpc_opt=module=github.com/bbengfort/oneof/pb messages.proto wrapper.proto
