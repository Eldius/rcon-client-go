
.tmp:
	mkdir -p .tmp


.tmp/server.jar: .tmp
	wget https://piston-data.mojang.com/v1/objects/f69c284232d7c7580bd89a5a4931c3581eae1378/server.jar -O .tmp/server.jar


build-testserver: .tmp/server.jar
	docker \
		buildx \
		build \
			--no-cache \
			--tag eldius/mine-rcon-server:latest \
				.
