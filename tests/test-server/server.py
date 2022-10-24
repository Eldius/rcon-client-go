import asyncio
from os import environ

from rcon_server.rcon_server import RCONServer
from rcon_server.rcon_packet import RCONPacket

class MyRCONServer(RCONServer):
    def __init__(self, bind=..., password=None):
        super().__init__(bind, password)

    def handle_execcommand(self, packet, connection):
        """
        Handles an EXECCOMMAND package. This command has to be implemented by
        a subclass of the RCONServer.
        :param packet: the packet containing the command
        :param connection: the RCONConnection which calls this method
        """
        print(f"command received: '{packet.body}' (type '{packet.type}')")
        response = RCONPacket(id=packet.id,type=RCONPacket.SERVERDATA_AUTH, body="default response")
        if packet.body == "help":
            print(f"help received: '{packet.body}' (type '{packet.type}')")
            response = RCONPacket(id=packet.id, type=RCONPacket.SERVERDATA_RESPONSE_VALUE, body="server help message")
        if packet.body == "test":
            print(f"test received: '{packet.body}' (type '{packet.type}')")
            response = RCONPacket(id=packet.id, type=RCONPacket.SERVERDATA_RESPONSE_VALUE, body="test command received")
        if packet.body == "echo":
            print(f"echo received: '{packet.body}' (type '{packet.type}')")
            response = RCONPacket(id=packet.id, type=RCONPacket.SERVERDATA_RESPONSE_VALUE, body=packet.body)
        connection.send_packet(response)

if __name__ == "__main__":
    port = 27015
    if environ.get("SERVER_PORT") is not None:
        port = environ.get("SERVER_PORT")
    server = MyRCONServer(bind=('0.0.0.0', port))
    if environ.get("SERVER_PASS") is not None:
        server.set_password(environ.get("SERVER_PASS"))
    asyncio.run(server.listen())
