# URI Shortener

A fast and lightweight URL shortener built with **Golang**, **Fiber**, and **Redis**.

Features

- Shorten URLs with customizable tail lengths.
- Track link metrics.
- High performance using Fiber and Redis.
- Easy setup with Docker and docker-compose.

## Prerequisites

- Go 1.20 or later
- Docker

## Run

1. Clone and Build
```bash
git clone https://github.com/pauuser/uri-shortener.git
cd uri-shortener
go build -o uri-shortener ./cmd
```

2. Set Up Redis

Start Redis using Docker:
```bash
docker-compose up -d
```

3. Create Configuration

Create a configs/settings.json file. Example:

```json
{
  "host": "0.0.0.0",
  "port": "5555",
  "redis": {
    "address": "127.0.0.1:6379",
    "database": 0,
    "password": "uriuser"
  },
  "link_configuration": {
    "tail_length": 6,
    "domain": "http://127.0.0.1:5555"
  }
}
```

4. Run the Application
```bash
./uri-shortener -path-to-config configs -config-file settings.json
```

## API

1. Create Short Link

Endpoint: `POST /`

Request:
```
{
    "url": "https://example.com"
}
```

Response:
```
{
    "short_url": "http://127.0.0.1:5555/abc123"
}
```

2. Retrieve Original Link

Endpoint: `GET /:tail`

Redirects to the original URL.

3. Retrieve Link Metrics

Endpoint: `GET /:tail/metrics`

Response:

```
{
    "clickCount": 42,
    "created_at_utc": "2025-01-01T12:00:00Z"
}
```
