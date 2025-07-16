package handlers

import (
	"golang-product-api-jwt/database"
	"golang-product-api-jwt/models"
	"github.com/gofiber/fiber/v2"
)

func GetProducts(c *fiber.Ctx) error {
	// Jalankan query SQL untuk ambil semua data dari tabel products
	rows, err := database.DB.Query("SELECT * FROM products")
	if err != nil {
		return err // Kirim error jika query gagal
	}
	defer rows.Close() // Tutup rows setelah selesai dibaca

	var products []models.Product // Slice untuk menampung semua produk

	// Loop hasil query, pindahkan ke slice products
	for rows.Next() {
		var p models.Product
		rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock) // Baca setiap baris dan masukkan ke struct
		products = append(products, p)
	}

	// Kembalikan data produk dalam bentuk JSON
	return c.JSON(products)
}


func CreateProduct(c *fiber.Ctx) error {
	var products []models.Product

	// Parse body JSON array ke slice of struct
	if err := c.BodyParser(&products); err != nil {
		return err
	}

	// Loop untuk masukkan tiap produk ke database
	for _, p := range products {
		_, err := database.DB.Exec("INSERT INTO products (name, price, stock) VALUES (?, ?, ?)", p.Name, p.Price, p.Stock)
		if err != nil {
			return err
		}
	}

	// Kembalikan response sukses
	return c.JSON(fiber.Map{"message": "Products created"})
}



func UpdateProduct(c *fiber.Ctx) error {
	id := c.Params("id") // Ambil parameter id dari URL (ex: /api/products/2)

	var p models.Product
	if err := c.BodyParser(&p); err != nil {
		return err
	}

	// Jalankan query UPDATE untuk mengubah data produk berdasarkan ID
	_, err := database.DB.Exec("UPDATE products SET name=?, price=?, stock=? WHERE id=?", p.Name, p.Price, p.Stock, id)
	if err != nil {
		return err
	}

	// Kembalikan response sukses
	return c.JSON(fiber.Map{"message": "Product updated"})
}

func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id") // Ambil parameter id dari URL

	// Jalankan query DELETE untuk hapus produk berdasarkan ID
	_, err := database.DB.Exec("DELETE FROM products WHERE id=?", id)
	if err != nil {
		return err
	}

	// Kembalikan response sukses
	return c.JSON(fiber.Map{"message": "Product deleted"})
}

