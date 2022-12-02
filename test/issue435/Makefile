regenerate:
	go install github.com/alkiranet/gogo-protobuf/protoc-gen-gogo
	protoc --gogo_out=. --proto_path=../../../../../:../../protobuf/:. *.proto
