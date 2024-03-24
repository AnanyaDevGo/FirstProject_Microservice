run:
	go run cmd/main.go

	protoc --go_out=. --go-grpc_out=. ./pkg/pb/product/product.proto

	 protoc --go_out=. --go-grpc_out=. ./pkg/pb/cart/cart.proto