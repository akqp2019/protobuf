regenerate:
	go install github.com/akqp2019/protobuf/protoc-gen-gogo
	protoc-min-version --version="3.0.0" --gogo_out=Mgoogle/protobuf/timestamp.proto=github.com/akqp2019/protobuf/types:. --proto_path=../../../../../:../../protobuf/:. issue260.proto
