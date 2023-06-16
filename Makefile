create-audio:
	protoc --proto_path=proto proto/audio-proto/*.proto --go_out=proto/audio-proto/
	protoc --proto_path=proto proto/audio-proto/*.proto --go-grpc_out=proto/audio-proto/

create-auth:
	protoc --proto_path=proto proto/auth-proto/*.proto --go_out=proto/auth-proto/
	protoc --proto_path=proto proto/auth-proto/*.proto --go-grpc_out=proto/auth-proto/

create-gateway:
	protoc --proto_path=proto proto/gateway-proto/*.proto --go_out=proto/gateway-proto/
	protoc --proto_path=proto proto/gateway-proto/*.proto --go-grpc_out=proto/gateway-proto/

create-api-gateway:
	protoc --proto_path=api-gateway api-gateway/gateway-proto/*.proto --go_out=api-gateway/gateway-proto/
	protoc --proto_path=api-gateway api-gateway/gateway-proto/*.proto --go-grpc_out=api-gateway/gateway-proto/

clean-audio:
	rm proto/audio-proto/*.go

clean-auth:
	rm proto/auth-proto/*.go

clean-gateway:
	rm proto/gateway-proto/*.go