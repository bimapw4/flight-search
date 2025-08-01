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


### 🔬 Observability
Prometheus metrics are available at:

```
http://localhost:3000/metrics
```

or
```
http://localhost:9090/query
```

### 📮 Postman Collection
For easier testing, you can use the provided Postman collection which includes all available API endpoints:

[👉 Access the collection here](https://drive.google.com/file/d/1vf1lPo5AL2klTde9tD8J7fUcvKhk9FGi/view?usp=sharing)

You can import it into Postman to try out endpoints such as:

* ```POST``` /api/flights/search

* ```GET``` /api/flights/search/:search_id/stream

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
3. For a more complete folder structure, you can see the readme in the main or provider folder.

