FROM python:3.11-alpine

WORKDIR /app

COPY server.py /app/server.py
COPY requirements.txt /app/requirements.txt

RUN pip install --upgrade pip && pip install -r requirements.txt

ENTRYPOINT [ "python", "/app/server.py" ]
