regenerate:
	go install github.com/alkiranet/gogo-protobuf/protoc-gen-gogo
	protoc --proto_path=.:$(GOPATH)/src/:$(GOPATH)/src/github.com/alkiranet/gogo-protobuf/protobuf/ \
	--gogo_out=. events.proto 
