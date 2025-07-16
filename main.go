package main

import (
	"golang-product-api-jwt/database"
	"golang-product-api-jwt/handlers"
	"golang-product-api-jwt/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/golang-jwt/jwt/v5"
	"time"
	"os"
)

func main() {
	godotenv.Load() // Load isi file .env ke environment Go
	database.Connect() // fungsi buat connect ke database

	app := fiber.New() // Inisialisasi Fiber App

	// Login route (tidak pakai middleware)
	app.Post("/login", handlers.Login)

	// Protected routes
	api := app.Group("/api", middleware.Protected())
	api.Get("/products", handlers.GetProducts)
	api.Post("/products", handlers.CreateProduct)
	api.Put("/products/:id", handlers.UpdateProduct)
	api.Delete("/products/:id", handlers.DeleteProduct)

	app.Listen(":3000") // App berjalan di port 3000

	// Endpoint GET /login — hanya untuk generate token dummy (tanpa cek database)
	app.Get("/login", func(c *fiber.Ctx) error {
	claims := jwt.MapClaims{ 	// Buat isi (claims) token

		"username": "admin",
		"exp":      time.Now().Add(time.Hour * 2).Unix(),
	}
	// Buat JWT token dengan algoritma HS256 dan isi (claims) di atas
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Tandatangani token pakai secret dari file .env (JWT_SECRET)
	s, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	return c.JSON(fiber.Map{"token": s})	// Kirim token ke client dalam bentuk JSON

})
}
