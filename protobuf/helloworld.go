//go:generate protoc -I=. -I=$GOPATH/src --go_out=plugins=grpc:. helloworld.proto

package protobuf
