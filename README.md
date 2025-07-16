---

````markdown
# 🚀 Product CRUD API with Go Fiber, MySQL & JWT

RESTful API sederhana menggunakan **Golang (Fiber)**, **MySQL**, dan **JWT Authentication**.  
Digunakan untuk mengelola data produk dengan fitur login, proteksi token, dan operasi CRUD.

---

## 📦 Fitur Utama

- ✅ Login pakai JWT Token
- ✅ Middleware untuk proteksi endpoint (JWT Auth)
- ✅ Operasi CRUD: Get, Create, Update, Delete Produk
- ✅ Password terenkripsi (bcrypt)
- ✅ Koneksi ke MySQL lokal (Laragon support)

---

## 🧰 Tech Stack

- Go + Fiber
- MySQL
- JWT (`github.com/golang-jwt/jwt/v4`)
- bcrypt

---

## ⚙️ Cara Menjalankan

### 1. Clone project dan masuk ke folder
```bash
git clone https://github.com/Kaesc/Golang-Product-CRUD-API.git
cd Golang-Product-CRUD-API
````

### 2. Buat file `.env`

```env
JWT_SECRET=your_secret_key_here
```

### 3. Jalankan

```bash
go mod tidy
go run main.go
```

---

## 🔐 Login

### Endpoint:

```
POST /login
```

### Body:

```json
{
  "username": "your_username",
  "password": "your_password"
}
```

### Response:

```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR..."
}
```

---

## 🔑 Gunakan Token JWT

Untuk semua request ke endpoint produk:

**Header:**

```
Authorization: Bearer {token}
Content-Type: application/json
```

---

## 🛒 CRUD Produk

| Method | Endpoint           | Keterangan         |
| ------ | ------------------ | ------------------ |
| GET    | /api/products      | Lihat semua produk |
| POST   | /api/products      | Tambah produk      |
| PUT    | /api/products/\:id | Ubah produk        |
| DELETE | /api/products/\:id | Hapus produk       |

### Contoh Body untuk POST / PUT:

```json
{
  "name": "Kaes Hoodie",
  "price": 150000,
  "stock": 20
}
```

---

## 📌 Catatan

* Token hanya berlaku 1 jam (bisa diperpanjang)

---

## 👨‍💻 Author

By [Kaes](https://github.com/Kaesc) – 2025
Latihan membuat REST API dengan Go & JWT 💻
