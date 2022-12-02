regenerate:
	protoc --proto_path=.:$(GOPATH)/src/:$(GOPATH)/src/github.com/alkiranet/gogo-protobuf/protobuf/ \
	--gogo_out=. issue312.proto 
	(cd events && make regenerate)
