regenerate:
	go install github.com/alkiranet/gogo-protobuf/protoc-min-version
	go install github.com/alkiranet/gogo-protobuf/protoc-gen-gogoslick
	protoc-min-version --version="3.0.0" --proto_path=.:$(GOPATH)/src/:$(GOPATH)/src/github.com/alkiranet/gogo-protobuf/protobuf/ \
	--gogoslick_out=Mgoogle/protobuf/timestamp.proto=github.com/alkiranet/gogo-protobuf/types:. timefail.proto 
