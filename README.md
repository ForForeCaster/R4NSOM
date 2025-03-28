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

Example of public key:
```
-----BEGIN RSA PUBLIC KEY-----
MIIBCgKCAQEAu2D49XqKWO/pqvwXGbmZSfUqNpPtyNZZ3QlRf+Q+JiCt3649/pIX
46w4fLdjvWEl8kF6DhGxjh43LLeaQJaVDL1V3Pvp8WuKHrHn7EzG9T/JB4GIFCgC
em7f7QZx2yCxBCDNCiPgi/YLzZAS6GNPvIbdbynmtuimJB9psr8nCVO5a8qhsS+x
u+W+24NIBeYxszx3yy6xeB6ysLNvgNDrvqY52cIAXusS+o0Qb9doYwmQVBLkU/K5
JjRqgjxs80mfSOJe7ZL/rR/rHEJoiqV7fN6X88xll4NoXdpl7S5MnSXMf0wtMZ3j
Jyd+z9VjWzaPdG61g9iLWnwe68UqV1ytCQIDAQAB
-----END RSA PUBLIC KEY-----
```

Example of private key:
```
-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAu2D49XqKWO/pqvwXGbmZSfUqNpPtyNZZ3QlRf+Q+JiCt3649
/pIX46w4fLdjvWEl8kF6DhGxjh43LLeaQJaVDL1V3Pvp8WuKHrHn7EzG9T/JB4GI
FCgCem7f7QZx2yCxBCDNCiPgi/YLzZAS6GNPvIbdbynmtuimJB9psr8nCVO5a8qh
sS+xu+W+24NIBeYxszx3yy6xeB6ysLNvgNDrvqY52cIAXusS+o0Qb9doYwmQVBLk
U/K5JjRqgjxs80mfSOJe7ZL/rR/rHEJoiqV7fN6X88xll4NoXdpl7S5MnSXMf0wt
MZ3jJyd+z9VjWzaPdG61g9iLWnwe68UqV1ytCQIDAQABAoIBACPlp8RFedggfCRC
370Qqq55vqDOlbUOZZBcLuYgqIxe4diSlbNdsyBtl7PC4WOAR0gCJbsoxhp7LOhO
80jw4DSc2CR6CV80EjWfsRX6vnnxApcCyHpOEO16LMGIpMrNWEQ5cXaCGYxOsoOq
2Ih3MCkmsCjR/u5nIx8T+oihKL4XHPwqHaDh4K/yYXpJvzNblKpFojdQ2ozTst9I
CTJho1PAu8Y/D2budwRpTOGVqmX9/maV/ZUh8p9k5CwRmws928s6yevnUFKeV5nT
thdGc5/qN0Nlu+K0n6thl8nRkC2m0U4SQIyGHaBnhmf0lmdYoxA2nGXXrEWBvz5l
M8SMM7UCgYEA8j6luBewVzLZYsE1snNLR4P7TpHyIeYJOXrIE92Hchiyx3UdsRTE
LWSYLE+RgWMxZtHsU0wB5c4VIWjLSgSBWD3xTR1aTxiM2uvqrmVZIWPoypVUumtT
ja7V60Xr205PWojlWdzz5jkn1g9++OqmgyDek6KXWM3nAlsFbc4KIP0CgYEAxgTG
eIcPT1MR3dheKVYCyRMN6UwYhPQ7Bw7s9DOhXIs+fHNq7KtJdB+sbgKF9R289tfQ
2q/bGE5XAn10JuBOuPgpMdolgitWxV/7fCkmrGoF+6lzh2EhCdcajtIcpRqjwb7J
f/dXXskPL1He1P+gnzKf7VipVNIkZCgO6YGcT/0CgYEAx1lervxvlZK26lntgoj2
rkeii6n3BIZ1mOO9uvtW90Hz/cgXQx6/wktCH1TrfPrQMLkYLgQim6bYJMJEaQkE
GEYBSWCabPTbFV2HpzSoAJ5jeDMjorWn8LSsPg0wZy0uiGaXtzfpTtBCRBpEdztD
BC4MfJoUynPHY1zcPtyL4VECgYBriApTR1ysAftQ0n+HyPpoQN2sFDO4d3xaf+Nz
VVpKDB+Zq3+kF/wigFS9xMcD7etAzL/REfBTia88Xe4mcmatka2lMcZuoqRTOKUz
rpEezdCD8mgXW0p6soHsjvMZQZctzjJUOHgMs/h4BOiSGGMcMHyigiQqOJekgBBh
A45j7QKBgDqkZZCVCjzful/oQKdtvetok4V0FfxrdDlsomBFvAOCpyFp0Krw3BjM
XtWzkNW6siHHoKddGQ7PZFWgnIKZuPYGV6aTwky3XrNaj7lhDr4WDn42L8V3rcTA
XN2eDR4tFZP27PJfI1KhRwfWuV0vZuo4bGZyrPQBpwXarkuLV9Rr
-----END RSA PRIVATE KEY-----
```

!!! You need to generate your personal keys !!!

2. Insert the content of public.key as value of publicKey in client.go
3. Build the executable file
```
go build client.go
```


## TODO 

- Implement decryptor