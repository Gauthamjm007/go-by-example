# Triple DES Algorithm

- Triple DES Algorithm is a triple cipher sequential Crypto Algorithm

- It takes 3 keys each of a min length of 8 - K1, K2 and K3

- It offers upto 2^112 security

### Encrypt Steps

- first you need to encrypt plain text using the key K1
- then decrypt using key K2
- then encrypt text obtained from above with K3

- This will provide a TripleDES encrypted text

### Decrypt Steps

- first you need decrypt using key K3
- Then encrypt using key K2
- Then encrypt using key K1

- This will provide the Plain originl Text

![Block Diagram](img/tripleDES.jpg)

#### :warning: &nbsp;&nbsp; WARNING

DES is cryptographically broken and should not be used for secure applications.
