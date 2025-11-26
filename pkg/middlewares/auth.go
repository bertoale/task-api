// Package middleware contains custom middleware functions
// Middleware adalah function yang dijalankan sebelum/sesudah handler
package middlewares

import (
	"strings"

	"rest-api/internal/auth"
	"rest-api/internal/database"
	"rest-api/pkg/config"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// Claims adalah struct untuk JWT payload
// Berisi user ID dan standard JWT claims (exp, iat, dll)
type Claims struct {
	ID uint `json:"id"` // User ID dari database
	jwt.RegisteredClaims
}

// Auth adalah middleware untuk autentikasi user
// Middleware ini akan:
// 1. Mengambil token dari Authorization header atau cookie
// 2. Memverifikasi dan parse JWT token
// 3. Mengambil user dari database berdasarkan ID di token
// 4. Menyimpan user object di context (c.Locals) untuk digunakan di handler
// Parameter:
//   - cfg: Config object yang berisi JWT secret
// Returns: Fiber handler function
// Usage: app.Get("/protected", middleware.Auth(cfg), handler)
func Auth(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Ambil token dari Authorization header (format: "Bearer <token>")
		token := c.Get("Authorization")
		if token != "" && strings.HasPrefix(token, "Bearer ") {
			token = strings.TrimPrefix(token, "Bearer ")
		} else {
			// Jika tidak ada di header, coba ambil dari cookie
			token = c.Cookies("token")
		}

		// Jika token tidak ditemukan di header maupun cookie
		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Akses ditolak. Token tidak ditemukan.",
			})
		}

		// Parse dan verify JWT token
		claims := &Claims{}
		tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(cfg.JWTSecret), nil // Secret key untuk verify signature
		})

		// Jika token invalid atau expired
		if err != nil || !tkn.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Token tidak valid atau kadaluarsa.",
			})
		}
		// Ambil user dari database berdasarkan ID di claims
		var user auth.User
		if err := database.DB.First(&user, claims.ID).Error; err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "User tidak ditemukan.",
			})
		}
		// Simpan user object di context untuk digunakan di handler
		// Cara akses di handler: user := c.Locals("user").(*auth.User)
		c.Locals("user", &user)
		return c.Next() // Lanjut ke handler berikutnya
	}
}
