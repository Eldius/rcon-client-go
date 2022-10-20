import asyncio
from os import environ

from rcon_server.rcon_server import RCONServer

if __name__ == "__main__":
    server = RCONServer()
    if environ.get("SERVER_PASS") is not None:
        server.set_password(environ.get("SERVER_PASS"))
    asyncio.run(server.listen())
