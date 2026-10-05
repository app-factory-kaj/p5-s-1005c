# greeter — PRD

## Problem Statement

Teams building services on this platform need a small, known-good reference
example to follow when standing up a new Go HTTP service. Without a minimal,
working example that already follows the organization's conventions
(`app-factory-kaj/e2e-reference`), each new service reinvents basic structure,
request handling, and response shape from scratch.

## Solution

Greeter is a small Go HTTP service that exposes a single greeting endpoint. It
takes a name via a query parameter and returns a JSON greeting, built and
structured following the conventions set out in `app-factory-kaj/e2e-reference`,
so it can serve as a lightweight, working reference for other services.

## Actors

- **Calling Service** — another component or client in the organization that
calls greeter's HTTP endpoint programmatically to obtain a greeting. There is
no human-facing UI; greeter is consumed only via its API. *assumed*

## User Stories

1. As a Calling Service, I want to call GET /hello with a name, so that I
 receive a JSON greeting addressed to that name.
2. As a Calling Service, I want a clear error response when I call GET /hello
 without a valid name, so that I can detect and handle the bad request
 instead of receiving a misleading greeting.

## Product Decisions

- Greeter is a backend-only HTTP service with no sign-in and no end-user UI;
it is called machine-to-machine, so no authentication flow applies at the
product level. *assumed*
- GET /hello requires the `name` query parameter; when it is missing or empty,
the service returns a 400 error with a validation message rather than a
generic or default greeting. *assumed*
- The service follows the structure and conventions demonstrated in
`app-factory-kaj/e2e-reference` rather than inventing its own.

## Out of Scope

- Any persistence, storage, or history of greetings requested.
- Any additional endpoints beyond the single greeting endpoint.
- Any human-facing UI or web application.
- Rate limiting, authentication, or authorization on the endpoint.

## Open Questions

None at this time.