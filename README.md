# ChainGuard — SIH26125 (Bharat Electronics Limited)

Decentralized identity, risk-adaptive access control, and NFT-based asset
ownership on Hyperledger Fabric, built for SIH problem statement 26125.

## Before you touch any code

Read these in order:
1. `docs/master-spec.md` — the full project spec, architecture, and pinned
   software versions everyone must use
2. `docs/team-plans/backend-team-individual-plan.md` or
   `docs/team-plans/frontend-team-individual-plan.md` — your day-by-day
   individual plan
3. `docs/api-contract.md` — REST endpoint contracts (backend writes this,
   frontend builds against it)

## Repo layout

```
chainguard/
├── fabric-network/     Hyperledger Fabric network config (from fabric-samples)
├── chaincode/          Go chaincode: identity, asset, access-control
├── backend/            Node.js API layer (Fabric Gateway SDK + risk engine)
├── frontend/           React app
├── docs/               Specs, architecture, security review, API contract
└── scripts/            Setup/deploy helper scripts
```

## Setup

See `scripts/setup.sh` for the guided local setup order. Do not run steps
out of order — chaincode must be deployed before the backend can connect,
and the backend must be running before frontend integration will work.

## Team

Backend/Fabric: Avikrit, Jeswin (Ayush supporting, Days 1-4)
Frontend: Abdullah, Safeer, Nazima (Ayush leading/monitoring)
