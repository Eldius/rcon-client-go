
TEST_SERVER_TAG := eldius/test-rcon-server:latest
MINE_SERVER_TAG := eldius/mine-rcon-server:latest
MINE_SERVER_PATH := $(PWD)/tests/mine-server
TEST_SERVER_PATH := $(PWD)/tests/test-server

build-testserver:
	cd $(TEST_SERVER_PATH) ; $(MAKE) -C $(TEST_SERVER_PATH) -f $(TEST_SERVER_PATH)/Makefile  build-testserver


build-mineserver:
	cd $(MINE_SERVER_PATH) ; $(MAKE) -C $(MINE_SERVER_PATH) -f $(MINE_SERVER_PATH)/Makefile  build-testserver SERVER_TAG=$(MINE_SERVER_TAG)


start-testserver: build-testserver
	-docker kill test
	docker run \
		--rm \
		-it \
		-d \
		-m 512m \
		--cpus=0.5 \
		-e "SERVER_PASS=StrongP@ss" \
		-p 27015:27015 \
		--name test \
		$(TEST_SERVER_TAG)

start-mineserver: build-mineserver
	-docker kill mine
	docker run \
		--rm \
		-it \
		-m 1g \
		--cpus=2.0 \
		-d \
		-p 27015:27015 \
		--name mine \
		$(MINE_SERVER_TAG)

run:
	-rm exec.log
	go run cmd/cli/main.go test --debug

console:
	-rm exec.log
	go run cmd/cli/main.go console -s localhost:27015 --debug

console-rpi:
	-rm exec.log
	go run cmd/cli/main.go console -s 192.168.100.183:25575 --debug

test:
	go test ./... -cover

vulncheck:
	govulncheck ./...

lint:
	golangci-lint run
