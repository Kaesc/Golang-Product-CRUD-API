# Golang Fiber API CRUD dengan JWT Authentication

REST API sederhana yang dibangun dengan Golang menggunakan framework Fiber dengan autentikasi JWT dan database MySQL.

## 🚀 Fitur

- **CRUD Operations** - Create, Read, Update, Delete
- **JWT Authentication** - Secure API dengan JSON Web Token
- **Bearer Token Authorization** - Semua endpoint CRUD memerlukan Bearer Token
- **MySQL Database** - Penyimpanan data menggunakan MySQL
- **Fiber Framework** - Fast HTTP web framework untuk Go

## 🛠️ Tech Stack

- **Language**: Go (Golang)
- **Framework**: Fiber v2
- **Database**: MySQL
- **Authentication**: JWT (JSON Web Token)
- **ORM**: [Sesuaikan dengan ORM yang digunakan - GORM/Database/SQL]

## 📋 Prerequisites

Pastikan Anda telah menginstall:
- Go 1.19 atau lebih baru
- MySQL Server
- Git

## 🔧 Installation

1. Clone repository ini:
```bash
git clone https://github.com/Kaesc/Golang-Product-CRUD-API.git
cd Golang-Product-CRUD-API
```

2. Install dependencies:
```bash
go mod download
```

3. Setup database MySQL dan buat database baru:
```sql
CREATE DATABASE [nama_database];
```

4. Copy file konfigurasi dan sesuaikan dengan environment Anda:
```bash
cp .env.example .env
```

5. Edit file `.env` dengan konfigurasi database dan JWT secret:
```env
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password
DB_NAME=your_database_name
JWT_SECRET=your_jwt_secret_key
```

6. Jalankan aplikasi:
```bash
go run main.go
```

Server akan berjalan di `http://localhost:3000`

## 📚 API Endpoints

### Authentication

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/login` | Login dan dapatkan JWT token |

### CRUD Operations (Memerlukan Bearer Token)

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/products` | Ambil semua data |
| GET | `/api/products/:id` | Ambil data berdasarkan ID |
| POST | `/api/products` | Buat data baru |
| PUT | `/api/products/:id` | Update data berdasarkan ID |
| DELETE | `/api/products/:id` | Hapus data berdasarkan ID |

## 🔐 Authentication

### Login
```bash
POST /api/auth/login
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "password123"
}
```

**Response:**
```json
{
  "status": "success",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

### Menggunakan Bearer Token
Untuk semua endpoint CRUD, sertakan Bearer Token di header:

```bash
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

## 📖 Contoh Penggunaan

### Login
```bash
curl -X POST http://localhost:3000/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "password123"
  }'
```

### 3. Akses Endpoint CRUD (dengan Bearer Token)
```bash
curl -X GET http://localhost:3000/api/products \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

## 🗂️ Struktur Project

```
.
├── main.go
├── go.mod
├── go.sum
├── database/
│   └── mysql.go
├── handlers/
│   ├── auth.go
│   └── product.go
├── middleware/
│   └── jwt.go
├── models/
│   ├── product.go
│   └── user.go
└── tools/
    └── hash.go
```

## 🔒 Security

- Semua endpoint CRUD dilindungi dengan JWT authentication
- Token harus disertakan dalam format Bearer Token
- Akses ditolak jika token tidak valid atau tidak ada
- Password di-hash menggunakan bcrypt

## 🚨 Error Handling

API akan mengembalikan error dalam format JSON:

```json
{
  "status": "error",
  "message": "Unauthorized - Token required"
}
```

## 📞 Contact

Kaesc - [https://github.com/Kaesc](https://github.com/Kaesc)

Project Link: [https://github.com/Kaesc/Golang-Product-CRUD-API](https://github.com/Kaesc/Golang-Product-CRUD-API)
