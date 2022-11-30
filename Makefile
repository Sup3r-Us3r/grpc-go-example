cert:
		cd grpc/cert; ./generate.sh; cd ../..

gen:
		protoc --proto_path=grpc grpc/protofile/*.proto --go_out=. --go-grpc_out=.
