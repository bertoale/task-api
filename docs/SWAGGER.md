# Swagger API Documentation

Dokumentasi ini menjelaskan cara generate, update, dan menulis dokumentasi endpoint API menggunakan Swagger di proyek Go Fiber ini.

## Cara Generate & Update Docs

1. **Install swag CLI** (sekali saja):

   ```bash
   go install github.com/swaggo/swag/cmd/swag@latest
   ```

   Pastikan `$GOPATH/bin` sudah ada di PATH.

2. **Generate/Update docs** (setiap ada perubahan endpoint/comment):
   Jalankan dari root folder `server/`:

   ```bash
   swag init -g cmd/main.go -o ./docs
   ```

   - `-g cmd/main.go` = entrypoint utama aplikasi
   - `-o ./docs` = output folder docs

3. **Akses Swagger UI**
   Jalankan server, lalu buka:
   [http://localhost:5000/swagger/index.html](http://localhost:5000/swagger/index.html)

## Contoh Penulisan Komentar Swagger

Tambahkan komentar di atas handler/controller (contoh untuk endpoint register user):

```go
// @Summary Register user
// @Description Register a new user
// @Tags Auth
// @Accept json
// @Produce json
// @Param data body auth.RegisterRequest true "User data"
// @Success 201 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /api/auth/register [post]
func (h *AuthController) Register(c *fiber.Ctx) error {
    // ...
}
```

### Contoh Lain (Endpoint Task)

```go
// @Summary Create new task
// @Description Buat task baru untuk user
// @Tags Tasks
// @Accept json
// @Produce json
// @Param data body task.CreateTaskRequest true "Task data"
// @Success 201 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /api/tasks [post]
func (h *TaskController) CreateTask(c *fiber.Ctx) error {
    // ...
}
```

> **Tips:**
>
> - Gunakan `@Tags` sesuai modul (Auth, User, Tasks, dll)
> - `@Param` untuk parameter body/query/path
> - `@Success`/`@Failure` untuk response
> - `@Router` untuk path dan method

## Referensi

- [swaggo/swag](https://github.com/swaggo/swag)
- [swaggo/fiber-swagger](https://github.com/swaggo/fiber-swagger)
- [Swagger Spec](https://swagger.io/specification/)
