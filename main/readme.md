# flight-api.

### 1. Project Structure
<pre>
├── bootstrap/                # App initialization (DB, DI, Migrations)
│   ├── db.go
│   ├── migrate.go
│   └── providers.go
│
├── internal/                 
│   ├── business/             # Business logic / usecases
│   ├── common/               # Common helpers (JWT, context, bcrypt, etc.)
│   ├── consts/                # Global constants
│   ├── entity/               # Request payloads / DTO (input layer)
│   ├── handlers/             # HTTP handlers (Fiber endpoints)
│   ├── middleware/           # Middleware (Audit log, Auth guard)
│   ├── presentations/        # DB models & API response structures
│   ├── provider/             # Dependency injection & service registry
│   ├── repositories/         # Data access layer (SQLX + PostgreSQL)
│   ├── response/             # API response wrapper (success / error)
│   ├── routes/                # HTTP route definitions (Fiber)
│   └── migrations/           # SQL migration scripts
│
├── pkg/                      
│   ├── databasex/            # Additional DB helper functions
│   └── meta/                 # Pagination, metadata utilities
│
├── .env                      # Environment variables
├── .env.example              # Sample environment file
├── docker-compose.yml        # Docker service setup
├── dockerfile                 # Dockerfile for app build
├── go.mod                     # Go modules
├── go.sum                     
├── main.go                   # Application entry point
└── readme.md                 
</pre>

### 👀 2. Features
> 
```
Database schema is auto-migrated on app startup (no need to run SQL manually)
```

### 3. Run the Project
Without Docker
```
go run main.go
```
##### or
with docker
```
docker-compose build --no-cache
docker-compose up
```

### 4. Migration
```
run the project go run main.go or use docker then the migration will run automatically
```


### 5. Technology Stack
* Golang (1.21+)

* Fiber (HTTP Framework)

* SQLX + PostgreSQL

* Redis

* Docker / Docker Compose

* Goose Migration


### 6. Env Example
```
APP_NAME = Cars Rent
PORT = 8083

DB_HOST = 
DB_USER = 
DB_PASSWORD = 
DB_NAME = 
DB_PORT = 

JWT_SECRET_KEY = 
JWT_LIFESPAN = 
```