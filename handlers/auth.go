package handlers

import (
	"database/sql"
	"golang-product-api-jwt/database"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)
func Login(c *fiber.Ctx) error {
	// Buat struct untuk menerima data dari request body
	type LoginInput struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var input LoginInput

	// Parse body JSON ke struct LoginInput
	if err := c.BodyParser(&input); err != nil {
		// Jika parsing gagal (misal format JSON salah), kirim error
		return err
	}

	// Siapkan variabel untuk menyimpan hasil query dari database
	var id int
	var hashedPassword string

	// Ambil data user dari database berdasarkan username
	err := database.DB.QueryRow(
		"SELECT id, password FROM users WHERE username = ?",
		input.Username,
	).Scan(&id, &hashedPassword)

	if err == sql.ErrNoRows {
		// Jika username tidak ditemukan di database
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "User not found"})
	} else if err != nil {
		// Error lainnya (koneksi, query error, dsb)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Server error"})
	}

	// Bandingkan password yang dikirim dengan hash password dari database
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(input.Password)); err != nil {
		// Jika password salah (tidak cocok dengan hash)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Wrong password"})
	}

	// Jika login berhasil → Buat JWT token
	claims := jwt.MapClaims{
		"user_id":  id,                                 // Simpan user_id di token
		"username": input.Username,                     // Simpan username
		"exp":      time.Now().Add(time.Hour * 1).Unix(), // Token expired dalam 1 jam
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // Buat token dengan claim dan HS256

	// Tandatangani token dengan JWT_SECRET dari file .env
	signedToken, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	// Kirim token ke client dalam response JSON
	return c.JSON(fiber.Map{"token": signedToken})
}
