FROM openjdk:17-alpine3.14

WORKDIR /app

COPY settings /app
COPY .tmp/server.jar /app/server.jar

ENTRYPOINT [ "java", "-Xmx600M", "-Xms600M", "-jar", "/app/server.jar", "nogui" ]
