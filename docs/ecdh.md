# Elliptic-Curve Diffie-Hellman (ECDH) in UDAL

## What is ECDH?

Elliptic-Curve Diffie-Hellman (ECDH) is an anonymous key agreement protocol that allows two parties, each having an elliptic-curve public-private key pair, to establish a shared secret over an insecure channel. 

This shared secret can be directly used as a key, or more commonly, it is passed through a Key Derivation Function (KDF) to derive symmetric keys that can then be used to encrypt subsequent communications using a symmetric-key cipher like AES-GCM or ChaCha20-Poly1305.

## What is an Elliptic Curve?

In mathematics, an elliptic curve is a specific type of algebraic curve. Visually, it looks a bit like a loop with an attached tail (or just a continuous loop, depending on the parameters).

For cryptography, we don't use a continuous curve. We use an elliptic curve mapped over a "finite field" (a grid of discrete points wrapped around a maximum prime number). This looks like a scatterplot of dots that still follows the underlying mathematical rules of the curve.

### The "Trapdoor" Function (Elliptic Curve Discrete Logarithm Problem)

The core property that makes elliptic curves useful for cryptography is a specific mathematical trick: **Point Addition**.

Imagine a game of cryptographic billiards:
1. You start at a known baseline point on the curve (the "Generator Point", standardized for everyone).
2. You pick a massive random number (your **Private Key**).
3. You "hit" the baseline point, causing it to bounce around the curve exactly the number of times specified by your Private Key.
4. The final resting place of the point is your **Public Key**.

The magic is that this is a **Trapdoor Function**:
- **Easy one way:** If I know the starting point and the number of times to bounce (Private Key), it is computationally very fast to calculate the final resting point (Public Key).
- **Hard the other way:** If I only know the starting point and the final resting point (Public Key), it is computationally *impossible* (it would take trillions of years for the world's supercomputers) to figure out how many times it bounced (Private Key) to get there.

### How it creates a Shared Secret (Diffie-Hellman)

1. Alice bounces the starting point `A` times (Private Key A) to get her Public Key A.
2. Bob bounces the starting point `B` times (Private Key B) to get his Public Key B.
3. They openly trade Public Keys.
4. Alice takes Bob's Public Key (which has already been bounced `B` times) and bounces it `A` more times. The final point has been bounced `A + B` times.
5. Bob takes Alice's Public Key (which has already been bounced `A` times) and bounces it `B` more times. The final point has been bounced `B + A` times.
6. Since `A + B = B + A`, they arrive at the exact same physical coordinate on the curve. 

This final coordinate is the **Shared Secret**. Anyone eavesdropping only saw the partially bounced points (Public Keys) and cannot figure out the final coordinate without knowing either `A` or `B`.

## How it works in UDAL

In the Universal Decentralized Anonymity Layer (UDAL), identity is decoupled from IP addresses. Instead, a cryptographic keypair represents a node's identity.

1. **Identity Generation**: When a UDAL node starts up, it generates an **X25519** keypair. 
   - The **Public Key** acts as the node's permanent address/identity in the network.
   - The **Private Key** is kept secret and used for authentication and encryption.

2. **Key Exchange**: When Node A wants to send an encrypted message to Node B:
   - Node A takes its own **Private Key** and Node B's **Public Key** to perform the ECDH mathematical operation. 
   - This operation yields a **Shared Secret**.

3. **Symmetric Encryption**: Node B performs the exact same ECDH operation using its own **Private Key** and Node A's **Public Key**. 
   - The mathematics of elliptic curves guarantee that Node B will arrive at the **exact same Shared Secret** as Node A.
   - Both nodes now share a secret key without ever having transmitted it across the network.
   - They can now use this derived shared secret to establish a fast, encrypted UDP tunnel (similar to how WireGuard operates).

## Why X25519?

We use the **Curve25519** elliptic curve (specifically the X25519 key exchange protocol) because:
- **Fast**: It is optimized for extremely fast, constant-time operations.
- **Secure**: It is highly resistant to side-channel attacks and does not rely on random number generators during the key exchange phase.
- **Standardized**: It is the modern standard for high-performance VPNs like WireGuard, as well as modern TLS 1.3 implementations.
