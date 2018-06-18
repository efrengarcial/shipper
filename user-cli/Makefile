build:
	docker build -t user-cli .

run:
	docker run --net="host" -e MICRO_REGISTRY=consul -e MICRO_REGISTRY_ADDRESS=localhost user-cli