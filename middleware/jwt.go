package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"strings"
)
// Protected adalah middleware untuk mengamankan route yang butuh autentikasi
func Protected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Ambil header "Authorization"
		auth := c.Get("Authorization")

		// Jika header kosong, tolak akses (Unauthorized)
		if auth == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Missing token",
			})
		}

		// Ambil token-nya saja, hapus prefix "Bearer "
		tokenStr := strings.Replace(auth, "Bearer ", "", 1)

		// Parse token JWT dengan secret dari .env
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			// Secret key yang digunakan untuk verifikasi token
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		// Jika token tidak valid atau error parsing, tolak akses
		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Invalid token",
			})
		}

		// Jika token valid, lanjut ke handler berikutnya
		return c.Next()
	}
}

