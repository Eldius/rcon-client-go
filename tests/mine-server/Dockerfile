FROM openjdk:17-alpine

WORKDIR /app

COPY settings /app
COPY .tmp/server.jar /app/server.jar

ENTRYPOINT [ "java", "-Xmx600M", "-Xms600M", "-jar", "/app/server.jar", "nogui" ]
