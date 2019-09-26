# grpc-service.felix

The default service template used for [felix](https://github.com/scottcrawford03/felix)

## Dependencies

Dependencies are included in `go.mod` by importing them inside a package that is never built, `tools/tools.go`. These dependencies should download automatically when `make build` is run. If they are not run automatically, install the following command line tools.

* [golint](https://github.com/golang/lint) `go get -u golang.org/x/lint/golint`
* [protoc-gen-grpc-gateway](github.com/grpc-ecosystem/grpc-gateway) `go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway`
* [protoc-gen-swagger](github.com/grpc-ecosystem/grpc-gateway) `go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger`
* [protoc-gen-go](github.com/golang/protobuf) `go get -u github.com/golang/protobuf/protoc-gen-go`
