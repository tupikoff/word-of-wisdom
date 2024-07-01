# Word of Wisdom

Design and implement “Word of Wisdom” tcp server.

• TCP server should be protected from DDOS attacks with the Proof of Work (https://en.wikipedia.org/wiki/Proof_of_work), the challenge-response protocol should be used.

• The choice of the POW algorithm should be explained.

• After Proof Of Work verification, server should send one of the quotes from “word of wisdom” book or any other collection of the quotes.

• Docker file should be provided both for the server and for the client that solves the POW challenge

---
## Challenge–response protocols with proof of work using the HashCash algorithm

1. Client requests access to the resource by sending "request".
2. Server responds with a challenge: "challenge|{{random_string:difficulty}}".
   1. random_string - a random string with a random length (20-30 symbols).
   2. difficulty - the number of leading zeroes for the HashCash algorithm (19-23 zeroes).
3. Client solves the challenge calculating HashCash, and answers with the solution: "response|{{hashcash}}".
4. Server checks:
   1. The random_string and difficulty in the solution hash are the same as those given to the client.
   2. The HashCash is valid.
   3. Stores the hash in storage and ensures it has never been used before.
5. Server grants access to the resource (and sends "words of wisdom" in the payload): "granted|{{payload}}", or closes the connection.

## Why HashCash

1. It was the first algorithm that I found on Wikipedia, and I like it.
2. As I know now, it was one of the foundational algorithms for POW and Bitcoins.

## TODO

1. Limit time for the solution.
2. Implement server graceful shutdown.
