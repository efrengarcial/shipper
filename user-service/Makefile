build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/efrengarcial/shipper/user-service \
		proto/auth/auth.proto
	docker build -t user-service .

run:
	docker run --net="host" -e MICRO_REGISTRY=consul -e MICRO_REGISTRY_ADDRESS=localhost  \
	    -e DB_NAME=postgres -e DB_HOST=localhost -e DB_PORT=5432 -e DB_USER=postgres -e DB_PASSWORD=postgres user-service