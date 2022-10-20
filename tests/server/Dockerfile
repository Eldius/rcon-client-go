FROM python:3.10-alpine3.16

WORKDIR /app

COPY server.py /app/server.py

RUN pip install --upgrade pip && pip install rcon-server

ENTRYPOINT [ "python", "/app/server.py" ]
