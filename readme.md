# ✈️ Flight Search System

An event-driven flight search system using Redis Streams, Fiber, and OpenTelemetry. Built to simulate integration with external airline providers, it delivers real-time results via Server-Sent Events (SSE).

## 📦 Tech Stack

- **Golang** + [Fiber](https://gofiber.io/)
- **Redis Streams** (XADD / XREADGROUP)
- **Prometheus**
- **Docker Compose**


## 🚀 Feature

- ✅ POST `/api/flights/search`  
  Send a search request to Redis Stream

- ✅ SSE `/api/flights/search/:search_id/stream`  
  Receive search results

- ✅ Graceful shutdown  
Redis, Fiber, and consumer will be shutdown clearly

## 🔄 How to Run

### 1. Clone & build
```bash
1. git clone https://github.com/bimapw4/flight-search.git
2. cd flight-search
3. cd main
    - enter the env file to connect to redis and the port
4. cd provider
    - enter the env file to connect to redis and the port
5. cd ..
    - enter the env file to connect to redis
6. docker compose up -d
```

### 2. Running Manual
```bash
1. git clone https://github.com/bimapw4/flight-search.git
2. cd flight-search
3. cd main
    - enter the env file to connect to redis and the port
4. go run main.go
5. cd ..
6. cd provider
    - enter the env file to connect to redis and the port
7. go run main.go
```

### 📚 Additional Note
1. Flight data read from ```sample.json``` in provider root
2. No database — all realtime via Redis

