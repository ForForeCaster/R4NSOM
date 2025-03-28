import socket
from Crypto.Cipher import AES
from Crypto.Random import get_random_bytes
import base64


def start_server():
    host = '0.0.0.0'
    port = 12345
    server_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    server_socket.bind((host, port))
    server_socket.listen(1)
    print("Server is listening...")

    while True:
        try:
            conn, addr = server_socket.accept()
            print(f"Connection from {addr} established.")

            aes_key = get_random_bytes(16)  # 128-bit key
            print("Generated AES key:", aes_key.hex())

            encoded = base64.b64encode(aes_key)
            conn.send(encoded)
            print(f"AES key '{encoded}' sent to client.")
            conn.close()
        except KeyboardInterrupt:
            print("Interrupted")
            break


if __name__ == "__main__":
    start_server()