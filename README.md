# Word of Wisdom

Design and implement “Word of Wisdom” tcp server.

• TCP server should be protected from DDOS attacks with the Proof of Work (https://en.wikipedia.org/wiki/Proof_of_work), the challenge-response protocol should be used.

• The choice of the POW algorithm should be explained.

• After Proof Of Work verification, server should send one of the quotes from “word of wisdom” book or any other collection of the quotes.

• Docker file should be provided both for the server and for the client that solves the POW challenge

---
## Challenge–response protocols

1. Client send request: "request"
2. Server challenge: "challenge|challenge_string:difficulty"
3. Client response with solution: "response|{{hashcash}}"
4. Server grant: "granted|{{payload}}"
or ignore

## TODO
1. server graceful shutdown
2. 