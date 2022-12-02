regenerate:
	go install github.com/alkiranet/gogo-protobuf/protoc-gen-gogo
	protoc-min-version --version="3.0.0" --gogo_out=Mgoogle/protobuf/timestamp.proto=github.com/alkiranet/gogo-protobuf/types:. --proto_path=../../../../../:../../protobuf/:. issue260.proto
