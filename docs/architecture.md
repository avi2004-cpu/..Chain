# Architecture

See the four-layer diagram in the master spec. Fill this in with any
architecture decisions made during the build that deviate from the
original plan — keep it current, don't let it go stale.

## Layers
1. Client layer — React frontend
2. Policy & risk layer — Node.js backend, risk engine
3. Blockchain layer — Hyperledger Fabric, 3 chaincodes + endorsement policy
4. Off-chain storage — encrypted document store
