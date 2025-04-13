# GOLANG-RESTFUL-API
Aplikasi ini merupakan implementasi RESTful API sederhana menggunakan bahasa Go (Golang) untuk manajemen data `Category`. API ini memiliki fitur `CRUD`, menggunakan database `MySQL`, dan menerapkan konsep Clean Architecture (layered).

## 🧱 Urutan Pembuatan Project

```markdown
1. model/domain        → Buat struktur tabel (Category)
2. repository          → Buat interface + query DB
3. service             → Implementasikan logic di atas repository
4. model/web           → Buat struct input/output API
5. controller          → Parsing request, panggil service, balikin response
6. middleware          → Tambahkan otorisasi (API Key)
7. helper              → Fungsi bantu response/error/tx
8. exception           → Buat NotFoundError atau BadRequest
9. app/router.go       → Mapping URL ke controller
10. app/database.go    → Setup koneksi database
11. main.go            → Jalankan server, inject semua dependency
12. test/category_test → Buat unit test / integration test
13. test.http          → Testing manual
```

## 📁 Struktur Folder dan Kegunaan

```markdown
golang-restful-api/
├── app/                # Inisialisasi router dan koneksi database
├── controller/         # Handler/controller yang menerima HTTP request
├── exception/          # Custom error handler dan middleware untuk error
├── helper/             # Fungsi-fungsi utilitas (JSON, TX, Error, dll)
├── middleware/         # Middleware (contoh: autentikasi API Key)
├── model/
│   ├── domain/         # Struktur data model untuk database
│   └── web/            # Struktur data untuk request & response API
├── repository/         # Interaksi langsung ke database (CRUD)
├── service/            # Logika bisnis aplikasi
├── test/               # Unit test untuk controller dan fitur API
├── main.go             # Entry point aplikasi
├── go.mod              # File module Go
├── go.sum              # Checksum dependency Go
└── test.http           # Contoh testing API menggunakan HTTP client
```

## ✅ Fitur

```markdown
1. Autentikasi menggunakan X-API-Key
2. API Endpoint:
    - GET /api/categories – Get all categories
    - GET /api/categories/{id} – Get category by ID
    - POST /api/categories – Create new category
    - PUT /api/categories/{id} – Update category
    - DELETE /api/categories/{id} – Delete category
3. Validasi input JSON
4. Struktur response konsisten (code, status, data)
5. Testing otomatis menggunakan httptest
```

## 💻 Cara Menjalankan Aplikasi
```markdown
1. Clone Repository
git clone https://github.com/rizkycahyono97/golang-restful-api.git
cd golang-restful-api

2. Setup Database (MySQL/MariaDB)
Buat database baru di MySQL:
    - CREATE DATABASE golang_api;
    - CREATE TABLE category(
      id int primary key auto_increment,
      name varchar(200) NOT NULL
      ) engine = innodb
    
3. Konfigurasi koneksi database
Di app/database.go, ubah sesuai kredensial lokal:
dsn := "root:password@tcp(localhost:3306)/golang_api"

4. Jalankan Aplikasi
go run main.go

5. Coba API dengan Postman atau REST Client
Contoh:
GET all categories
GET http://localhost:8080/api/categories
X-API-Key: RAHASIA
Accept: application/json
```

## 🧪 Cara Testing
```markdown
1. Jalankan semua test:
go test ./test

Jika muncul error dependency seperti:
missing go.sum entry for module ...

Jalankan:
go mod tidy

2. Test dengan HTTP
    test.http
```

## 📦 Dependency
```bash
github.com/go-sql-driver/mysql – Driver MySQL
github.com/gorilla/mux – Router HTTP
github.com/stretchr/testify – Library testing (assert)
context – Built-in Go untuk handling request context
```

### 🔐 API Key
```markdown
Untuk mengakses semua endpoint, tambahkan header:
X-API-Key: RAHASIA
```

### 📌 Catatan
```bash
Struktur response API menggunakan format:


{
  "code": 200,
  "status": "OK",
  "data": [...]
}
```

👨‍💻 Author
*Rizky Cahyono Putra* 
Built with ❤️ using Golang

---