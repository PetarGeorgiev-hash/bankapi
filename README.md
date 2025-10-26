# 🏦 bankapi

**bankapi** is a powerful backend API built with **Go**, designed for a simple yet scalable banking system.  
It provides functionality for managing users, accounts, and money transfers, with both **REST** and **gRPC** interfaces, **PostgreSQL** integration, **JWT authentication**, and **Docker** support.

---

## 🚀 Features

- 🧩 Clean and modular architecture in **Go**
- 🗃️ **PostgreSQL** database with **SQLC** for type-safe queries
- 🌐 **REST** and **gRPC** APIs
- 🔒 Secure **JWT authentication** and token management
- 📬 Email service for notifications
- 🧰 **Docker Compose** for easy development setup
- 🧪 Ready for **unit and integration testing**
- 🏗️ Simple to extend with new banking features

---

## 🛠️ Getting Started

### 1️⃣ Clone the repository

````bash
git clone https://github.com/PetarGeorgiev-hash/bankapi.git
cd bankapi

---

### 2️⃣ Configure environment variables

Create a `.env` file in the root directory and add the following configuration values:

```env
DB_DRIVER=postgres
DB_SOURCE=postgresql://user:password@localhost:5432/bank?sslmode=disable
SERVER_ADDRESS=0.0.0.0:8080
GRPC_SERVER_ADDRESS=0.0.0.0:9090
TOKEN_SYMMETRIC_KEY=your_jwt_secret_key
ACCESS_TOKEN_DURATION=15m
SMTP_HOST=smtp.example.com
SMTP_USER=example@example.com
SMTP_PASSWORD=yourpassword
SMTP_PORT=587


Make sure your PostgreSQL database is running and accessible.

🐳 Option 1: Run with Docker
docker-compose up --build


💻Option 2: Run locally
make server or go run main.go

To Run tests
To ensure everything is working correctly, run:
make test


📜 API Documentation
REST API

Once the server is running, you can access the Swagger documentation at:

http://localhost:8080/swagger/index.html
````
