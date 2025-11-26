# Go Task API – Backend

## 1. Project Overview

Go Task API adalah backend RESTful API untuk aplikasi manajemen tugas (todo) dengan autentikasi user, task management, dan arsitektur clean code berbasis vertical layer. Backend ini dibangun menggunakan Go, Fiber, dan GORM (ORM untuk MySQL).

**Teknologi utama:**

- Go (Golang)
- Fiber (web framework)
- GORM (ORM)
- MySQL
- JWT Authentication

---

## 2. Features

- User registration & login (JWT authentication)
- CRUD tasks (create, read, update, delete)
- User profile management
- Middleware: JWT auth, error handler, CORS, logger
- Clean architecture (vertical layer)
- Swagger API documentation

---

## 3. Tech Stack

- **Backend Framework:** [Go Fiber](https://gofiber.io/)
- **ORM/Database:** [GORM](https://gorm.io/) + MySQL
- **Authentication:** JWT (JSON Web Token)
- **Other Utilities:**
  - [Swaggo](https://github.com/swaggo/swag) (Swagger docs)
  - [godotenv](https://github.com/joho/godotenv) (env loader)
  - [Gofiber Middleware](https://docs.gofiber.io/api/middleware/)

---

## 4. Project Structure

```
server/
│
├── cmd/                # Entry point aplikasi (main.go)
├── internal/
│   ├── auth/           # Modul autentikasi (model, repository, service, controller, route)
│   ├── user/           # Modul user/profile (model, repository, service, controller, route)
│   ├── task/           # Modul task/todo (model, repository, service, controller, route)
│   ├── database/       # Koneksi & migrasi database
│   ├── routes/         # Setup routing utama (vertical_routes.go)
│
├── pkg/
│   ├── config/         # Konfigurasi aplikasi (env, dsb)
│   ├── response/       # Response helper (Success/Error)
│   ├── middlewares/    # Middleware global (auth, error handler)
│
├── docs/               # Dokumentasi Swagger (auto-generated)
├── .env                # Environment variables
├── go.mod              # Go modules
└── README.md           # Dokumentasi ini
```

**Penjelasan:**

- `internal/` berisi logic utama per fitur (auth, user, task)
- `pkg/` berisi helper, config, dan middleware global
- `docs/` berisi dokumentasi Swagger
- `cmd/` entry point aplikasi

---

## 5. Installation

```bash
# Clone repository
git clone <repo-url>
cd server

# Install dependencies
go mod tidy
```

---

## 6. Environment Variables

Buat file `.env` di root folder. Berikut variabel yang diperlukan:

| Variable       | Contoh Nilai              | Keterangan                 |
| -------------- | ------------------------- | -------------------------- |
| DB_HOST        | localhost                 | Host database MySQL        |
| DB_PORT        | 3306                      | Port database              |
| DB_USER        | root                      | Username database          |
| DB_PASSWORD    | (isi password)            | Password database          |
| DB_NAME        | libgo                     | Nama database              |
| DB_SSLMODE     | disable                   | SSL mode (disable/default) |
| JWT_SECRET     | your_super_secret_jwt_key | Secret key JWT             |
| JWT_EXPIRES_IN | 168h                      | Expiry JWT (contoh: 168h)  |
| PORT           | 5000                      | Port aplikasi              |
| NODE_ENV       | development               | Mode aplikasi              |
| CORS_ORIGIN    | http://localhost:3000     | Origin frontend            |

**Contoh .env:**

```env
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=
DB_NAME=libgo
DB_SSLMODE=disable

JWT_SECRET=your_super_secret_jwt_key
JWT_EXPIRES_IN=168h

PORT=5000
NODE_ENV=development
CORS_ORIGIN=http://localhost:3000
```

---

## 7. Database Setup

- **Migrasi otomatis** dijalankan saat server start (lihat `main.go`).
- Untuk migrasi manual, gunakan GORM migrasi di kode.
- **Seed data:** (tambahkan script seed jika diperlukan, default tidak ada).

---

## 8. Running the App

### Development

```bash
go run ./cmd/main.go
```

### Production

```bash
go build -o app ./cmd/main.go
./app
```

### NPM Scripts

> Tidak menggunakan npm, semua perintah via Go CLI.

---

## 9. API Documentation

Swagger UI tersedia di:  
[http://localhost:5000/swagger/index.html](http://localhost:5000/swagger/index.html)

### Endpoint Utama

#### Auth

- `POST /api/auth/register` – Register user baru
- `POST /api/auth/login` – Login, dapatkan JWT

#### User

- `GET /api/users/profile` – Lihat profil user (auth)
- `PUT /api/users/:id` – Update profil user (auth)

#### Tasks

- `POST /api/tasks` – Buat task (auth)
- `GET /api/tasks` – List task user (auth)
- `GET /api/tasks/:id` – Detail task (auth)
- `PUT /api/tasks/:id` – Update task (auth)
- `DELETE /api/tasks/:id` – Hapus task (auth)

### Contoh Request Register

```json
{
  "name": "Budi",
  "email": "budi@mail.com",
  "password": "passwordku"
}
```

### Contoh Response Sukses

```json
{
  "success": true,
  "message": "Register success",
  "data": {
    "id": 1,
    "name": "Budi",
    "email": "budi@mail.com"
  }
}
```

---

## 10. Authentication

- Menggunakan JWT (JSON Web Token).
- Setelah login, user mendapat token yang dikirim di header `Authorization: Bearer <token>`.
- Middleware akan memproteksi route yang membutuhkan autentikasi.

---

## 11. Error Handling

Format error global:

```json
{
  "success": false,
  "message": "error message"
}
```

Contoh error:

```json
{
  "success": false,
  "message": "Invalid email or password"
}
```

---

## 12. Testing

> (Opsional, tambahkan jika ada unit test)

```bash
go test ./...
```

---

## 13. Deployment

- **Docker:**  
  Tambahkan Dockerfile dan docker-compose sesuai kebutuhan.
- **Railway/Render/VPS:**  
  Deploy dengan build Go binary, set environment variable sesuai server.
- **Production:**  
  Jalankan `go build` lalu jalankan binary.

---

## 14. License

MIT

---
