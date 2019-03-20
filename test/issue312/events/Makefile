regenerate:
	go install github.com/akqp2019/protobuf/protoc-gen-gogo
	protoc --proto_path=.:$(GOPATH)/src/:$(GOPATH)/src/github.com/akqp2019/protobuf/protobuf/ \
	--gogo_out=. events.proto 
