## Intro

Simple Ransomware. Consist on server and client parts. The server generate the encryption key and send it to the client. Client using the key for encryption.

⚠️ WARNING: This software is made just for study purposes. ⚠️ WARNING: If you want to run it locally for tests, take care of what directories you decide to encrypt. The author doesn't take responsibility for any illegal use of the code by 3rd parties.

## Install

### Server

```bash
python -m venv venv
source ./venv/bin/activate
pip install pycryptodome
python server.py
```

### Client

1. Generate RSA pair
```bash
openssl genrsa -out private.key -pkeyopt rsa_keygen_bits:2048
openssl rsa -in private.key -pubout -out public.key
```

!!! You need to generate your personal keys !!!

2. Insert the content of public.key as value of publicKey in client.go
3. Build the executable file
```
go build client.go
```


## TODO 

- Implement decryptor
