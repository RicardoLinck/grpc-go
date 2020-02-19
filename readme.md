# gRPC in Go
Exercises from https://www.udemy.com/course/grpc-golang.

# Build protobuf with protoc

```bash
brew install protobuf
go get -u google.golang.org/grpc 
go get -u github.com/golang/protobuf/protoc-gen-go 
```

Make sure to add `protoc-gen-go` to your PATH.

## Greet
```bash
protoc greet/greetpb/greet.proto --go_out=plugins=grpc:. 
```

## Calculator
```bash
protoc calculator/calculatorpb/calculator.proto --go_out=plugins=grpc:.
```
