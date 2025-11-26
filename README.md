# Go Todo API (Vertical Layer Architecture)

A simple RESTful API for a Todo application built with Go, Fiber, and GORM (MySQL). This project implements user authentication (JWT), task management, and follows clean architecture principles (repository, service, controller).

## Features

- User registration & login (JWT authentication)
- CRUD tasks (create, read, update, delete)
- Each task belongs to a user
- Secure endpoints (protected with JWT)
- Clean code structure (separation of concerns)
- Error handling & validation

## Tech Stack

- [Go](https://golang.org/)
- [Fiber](https://gofiber.io/) (web framework)
- [GORM](https://gorm.io/) (ORM for MySQL)
- [MySQL](https://www.mysql.com/) (database)
- JWT for authentication

## Struktur Proyek

```
internal/
  auth/         # Modul autentikasi (model, repository, service, controller, route)
  user/         # Modul user/profile (model, repository, service, controller, route)
  task/         # Modul task/todo (model, repository, service, controller, route)
  middlewares/  # Middleware global (auth, error handler)
  database/     # Koneksi & migrasi database
  routes/       # Setup routing utama (vertical_routes.go)
pkg/
  config/       # Konfigurasi aplikasi (env, dsb)
  response/     # Response helper (Success/Error)
  middlewares/  # Middleware global (auth, error handler)

client/         # Frontend Next.js (app router, shadcn/ui)
cmd/            # Entry point aplikasi (main.go)
```

## Arsitektur

- **Vertical Layer**: Setiap fitur (auth, user, task) punya folder sendiri berisi model, repository, service, controller, dan route.
- **Tidak ada folder controllers/services/repositories global** (horizontal layer sudah tidak dipakai).
- **Middleware** hanya untuk validasi token dan error handling.
- **Generate Token** sekarang ada di `internal/auth/service.go`.

## Getting Started

### Prerequisites

- Go 1.18+
- MySQL server

### Installation

1. Clone the repo:
   ```bash
   git clone <repo-url>
   cd go-todo
   ```
2. **Copy `.env.example` ke `.env`** dan sesuaikan konfigurasi database.
3. Install dependencies:
   ```bash
   go mod tidy
   ```
4. **Jalankan migrasi & server**:
   ```bash
   go run ./cmd/main.go
   ```

## API Endpoints

### Auth

- `POST /api/auth/register` — Register new user
- `POST /api/auth/login` — Login, returns JWT token

### User

- `GET /api/users/profile` — Lihat profil user (auth)
- `PUT /api/users/:id` — Update profil user (auth)

### Tasks

- `POST /api/tasks` — Buat task (auth)
- `GET /api/tasks` — List task user (auth)
- `GET /api/tasks/:id` — Detail task (auth)
- `PUT /api/tasks/:id` — Update task (auth)
- `DELETE /api/tasks/:id` — Hapus task (auth)

## License

MIT

---

**Happy coding!** 🚀
