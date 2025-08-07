# Note: should use protoc-gen-grpc-gateway v2
servicepb:
	protoc --proto_path=./client/proto --go_out ./client/serviceapi --go_opt=module=github.com/brevis-network/prover-network-gateway/client/serviceapi \
	--go-grpc_out=./client/serviceapi --go-grpc_opt=require_unimplemented_servers=false,module=github.com/brevis-network/prover-network-gateway/client/serviceapi \
	--grpc-gateway_out ./client/serviceapi --grpc-gateway_opt=module=github.com/brevis-network/prover-network-gateway/client/serviceapi \
	./client/proto/prover_network.proto