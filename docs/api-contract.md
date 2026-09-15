# API Contract

Backend fills this in on Day 6 and treats it as a promise to frontend —
changing a shape after frontend has built against it should be rare and
communicated immediately.

## POST /identity/register
Request: { }
Response: { }

## GET /identity/:address
Request: —
Response: { }

## POST /assets/mint
Request: { }
Response: { }

## GET /assets/:id
Request: —
Response: { }

## POST /access/request
Request: { }
Response: { }

## POST /access/decision
Request: { }
Response: { }

## GET /access/logs
Request: —
Response: { }
