# GoLang Prep 🚀

This repository tracks my daily progress as I learn the Go programming language. Each folder contains code and exercises for that day. 

## **1. Go Language Fundamentals**

Before backend work, you must be very strong in:

* Variables, constants, data types
* Slices, arrays, maps
* Structs
* Pointers
* Functions, methods, interfaces
* Error handling (very important)
* Goroutines & Channels (Concurrency)

---

# ✅ **2. Go Project Structure**

Backend projects in Go follow clean structures like:

```
/cmd
/internal
/pkg
/config
/api
/models
/controllers
/services
/repository
```

Or simpler versions for small projects:

```
main.go
handlers/
models/
db/
utils/
```

---

# ✅ **3. REST API DEVELOPMENT**

You must know:

* How to create HTTP servers (net/http)
* Routing (gorilla/mux, chi)
* Request parsing (JSON → struct)
* Response formatting (struct → JSON)
* Middleware (auth, logging, validation)
* Handling HTTP methods (GET, POST, PUT, DELETE)

---

# ✅ **4. Dependency Management**

* Go Modules: `go mod init`, `go mod tidy`

---

# ✅ **5. Database Integration**

Know how to work with:

### **SQL Databases**

* PostgreSQL (best)
* MySQL

Using:

* `database/sql`
* ORMs like:

  * GORM
  * sqlx

Important concepts:

* Migrations (go-migrate)
* Connection pooling
* Transactions
* Joins, indexes

---

# ✅ **6. Authentication & Authorization**

You must know how to build secure auth:

### **Types**

* JWT tokens
* Session-based auth
* OAuth2 (Google, GitHub login)
* Role-based access control (RBAC)

---

# ✅ **7. Concurrency (Big advantage of Go)**

Master:

* Goroutines
* Channels
* Buffered channels
* Mutex & WaitGroup
* worker pools
* concurrency-safe data

---

# ✅ **8. Caching**

* Using Redis
* In-memory caching (map + mutex)

---

# ✅ **9. Clean Architecture & Design Patterns**

Learn how to structure large apps:

* Clean Architecture
* Repository Pattern
* Service Layer Pattern
* Middleware Pattern
* Dependency Injection (simple in Go)
* SOLID principles (Go style)

---

# ✅ **10. Logging & Monitoring**

* zerolog
* logrus
* slog (Go 1.21+)
* Request logging middleware
* Panic recovery middleware

Monitoring:

* Prometheus & Grafana
* Health checks `/health`

---

# ✅ **11. Testing**

Backend developers must write tests:

* Unit tests (`*_test.go`)
* Integration tests
* Mocking (testify/mock)
* Benchmarks (`go test -bench`)

---

# ✅ **12. Docker & Deployment**

Know how to:

* Write a Dockerfile for Go
* Build minimal Go images (scratch/alpine)
* Docker Compose for DB + app

Deploying on:

* AWS EC2
* Railway.app
* Render.com
* Kubernetes basics

---

# ✅ **13. Message Brokers (Optional but important)**

For scalable backend systems:

* RabbitMQ
* Kafka
* NATS
* SQS (AWS)

---

# ✅ **14. gRPC (Optional but modern)**

Microservices communication:

* proto files
* generating Go code with protoc
* gRPC services & streaming

---

# 🎯 **Most Important (Top 5) to Get a Job**

If you're preparing for backend jobs, focus on:

1. **REST APIs**
2. **GORM + PostgreSQL**
3. **Concurrency (goroutines, channels)**
4. **Authentication (JWT)**
5. **Clean Architecture**

---