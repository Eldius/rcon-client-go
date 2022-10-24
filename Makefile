
TEST_SERVER_TAG := test-rcon-server
MINE_SERVER_TAG := mine-rcon-server
MINE_SERVER_PATH := $(PWD)/tests/mine-server
TEST_SERVER_PATH := $(PWD)/tests/test-server

build-testserver:
	cd $(TEST_SERVER_PATH) ; $(MAKE) -C $(TEST_SERVER_PATH) -f $(TEST_SERVER_PATH)/Makefile  build-testserver SERVER_TAG=$(TEST_SERVER_TAG)


build-mineserver:
	cd $(MINE_SERVER_PATH) ; $(MAKE) -C $(MINE_SERVER_PATH) -f $(MINE_SERVER_PATH)/Makefile  build-testserver SERVER_TAG=$(MINE_SERVER_TAG)


start-testserver: build-testserver
	-docker kill test
	docker run \
		--rm \
		-it \
		-d \
		-e "SERVER_PASS=StrongP@ss" \
		-p 27015:27015 \
		--name test \
		$(TEST_SERVER_TAG)

start-mineserver: build-mineserver
	-docker kill mine
	docker run \
		--rm \
		-it \
		-d \
		-p 27015:27015 \
		--name mine \
		$(MINE_SERVER_TAG)

run:
	-rm exec.log
	go run cmd/cli/main.go test

