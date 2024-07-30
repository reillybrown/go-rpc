# Go RPC

This project serves as an easy-to-use boilerplate for generating RPC services.

- [x] simple gRPC example
- [x] centralized proto files
- [x] uses TCP
- [x] add Makefile
- [ ] auto compiles any .proto changes
- [ ] cross-language examples
- [ ] builds a server cli
- [ ] builds a client cli
- [ ] simple web ui

# Quickstart
1. run TCP server
```sh
go run server/main.go
```

2. in separate shell, run client
```sh
go run client/main.go [TEST_MESSAGE]
```
> the client is transient, and each run will add a message [TEST_MESSAGE] to the list