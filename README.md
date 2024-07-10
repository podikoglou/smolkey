# Smolkey
Simple, Fast, In-Memory Key-Value store.

## Features
- In-memory store for *very* quick access.
- Basic Operations: PUT, GET, ~~DELETE~~
- Exposed via REST API
- ~~REPL for interacting with the store~~
- ~~Metrics, accessible via the REPL~~

## Frequently Asked Questions
- [Why?](#why)
- [What about security?](#what-about-security)

### Why?
It's a toy project.

### What about security?
Smolkey follows a philosophy similar to that of Redis. We trust you to run
Smolkey under a firewall, so we really don't need to implement authentication
in the application level. (if you really want authentication, you can
theoretically proxy it through something like
[Caddy](https://caddyserver.com/).)
