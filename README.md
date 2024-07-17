# Smolkey
Simple, Fast, In-Memory Key-Value store.

## Features
- In-memory store for *very* quick access
- Basic Operations: PUT, GET, ~~DELETE~~
- Exposed via REST API

*Potential Future Features*:
- REPL for interacting with the store
- Metrics, accessible via the REPL

## Benchmarks
All benchmarks were performed on an M1 Mac using `wrk` with 6 threads and 60 connections.

### PUT
**Writing Random Keys with Random Values:**
```
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    11.79ms   32.21ms 251.35ms   89.68%
    Req/Sec    36.05k    28.12k  159.33k    76.73%
  2030546 requests in 10.10s, 247.87MB read
Requests/sec: 201094.97
Transfer/sec:     24.55MB
```

### GET
**Reading Random Keys:**
```
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    12.28ms   32.68ms 269.33ms   89.47%
    Req/Sec    37.82k    30.49k  179.97k    75.25%
  2103317 requests in 10.05s, 243.71MB read
  Non-2xx or 3xx responses: 59431
Requests/sec: 209377.07
Transfer/sec:     24.26MB
```

**Reading Specific Key:**
```
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    11.59ms   32.35ms 256.85ms   90.28%
    Req/Sec    46.51k    40.60k  209.97k    74.56%
  2626620 requests in 10.09s, 298.09MB read
Requests/sec: 260349.99
Transfer/sec:     29.55MB
```

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
