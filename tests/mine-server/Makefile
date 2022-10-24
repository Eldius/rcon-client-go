
build-testserver: .tmp/server.jar
	docker build -t $(SERVER_TAG) .


tests/server/.tmp:
	mkdir -p .tmp

.tmp/server.jar: tests/server/.tmp
	wget https://piston-data.mojang.com/v1/objects/f69c284232d7c7580bd89a5a4931c3581eae1378/server.jar -O .tmp/server.jar
