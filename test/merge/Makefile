regenerate:
	go install github.com/akqp2019/protobuf/protoc-min-version
	go install github.com/akqp2019/protobuf/protoc-gen-gogo
	protoc-min-version --version="3.0.0" --gogo_out=. \
	--proto_path=../../../../../:../../protobuf/:. merge.proto
