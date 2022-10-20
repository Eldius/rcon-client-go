
build-testserver:
	cd tests/server ; docker build -t rcon-server . 


start-testserver: build-testserver
	-docker kill test
	docker run \
		--rm \
		-it \
		-d \
		-p 27015:27015 \
		-e "SERVER_PASS=StrongP@ss" \
		--name test \
		rcon-server

run:
	go run cmd/cli/main.go test
