.PHONY: proto
proto:
	protoc -I pb pb/pubsub.proto --go_out=plugins=grpc:pb