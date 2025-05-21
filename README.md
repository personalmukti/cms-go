# CMS-Go

📘 **CMS-Go** adalah sistem manajemen konten (Content Management System) berbasis backend **Go (Golang)** yang dikembangkan sebagai alternatif ringan dan fleksibel dari platform seperti WordPress atau Blogspot. Proyek ini bersifat modular, dapat digunakan secara lokal tanpa layanan berbayar, dan siap diintegrasikan dengan frontend HTML, Vue, atau React.

---

## 🚀 Status Pengembangan

> ✅ **Backend inti telah selesai 100%**  
> Semua fitur dasar CMS sudah berfungsi penuh, termasuk manajemen user, role, artikel, kategori, tag, halaman statis, autentikasi JWT, serta kontrol akses berbasis role.

---

## 🔧 Teknologi yang Digunakan

| Teknologi        | Deskripsi                                          |
|------------------|----------------------------------------------------|
| [Go](https://go.dev/)             | Bahasa pemrograman utama                         |
| [Echo](https://echo.labstack.com/)           | Web framework untuk routing & middleware        |
| [PostgreSQL](https://www.postgresql.org/)     | Database utama via GORM                         |
| [GORM](https://gorm.io/)                     | ORM untuk pengelolaan database relasional       |
| [JWT](https://jwt.io/)                       | Autentikasi & otorisasi berbasis token          |
| Multipart/FormData         | Upload gambar ke local storage (`/uploads`)       |
| .env                        | Konfigurasi fleksibel berbasis file lingkungan    |

> 📦 Docker belum digunakan (pengembangan lokal manual via Laragon/VSCode)

---

## 📂 Struktur Proyek

```bash
cms-go/
├── config/               # Konfigurasi environment (.env)
├── controllers/          # Handler semua endpoint API
├── database/             # Koneksi & migrasi database
├── middleware/           # Autentikasi & Role Middleware
├── models/               # Struktur tabel & relasi database
├── response/             # Fungsi standar respon JSON
├── routes/               # Routing modular per fitur
├── seeders/              # Seeder data default (role, user, artikel)
├── utils/                # Helper (hashing, slug generator, dll)
├── uploads/              # Folder gambar artikel (local)
├── .env                  # File konfigurasi
├── go.mod
└── main.go               # Entry point aplikasi
```

---

## ⚙️ Fitur Utama

| Fitur                          | Status |
|--------------------------------|--------|
| ✅ Manajemen Artikel (CRUD + Gambar + Slug)         | ✔️ |
| ✅ Manajemen Kategori & Tag (Relasional)           | ✔️ |
| ✅ Halaman Statis Dinamis (About, Contact, dll)     | ✔️ |
| ✅ Upload Gambar via FormData (local folder)        | ✔️ |
| ✅ Autentikasi & Registrasi dengan JWT              | ✔️ |
| ✅ Role-based Access Control (Admin, Editor, Operator) | ✔️ |
| ✅ User Manager (List & Update Role User)           | ✔️ |
| ✅ Role Manager (CRUD Role Dinamis)                 | ✔️ |
| ✅ Token Refresh Endpoint                           | ✔️ |
| ✅ Public API: Artikel by Slug, Search, Pagination  | ✔️ |

---

## 🛠️ Fitur Opsional yang Akan Ditambahkan

| Fitur Opsional               | Status   |
|-----------------------------|----------|
| 🌐 Integrasi MinIO / CDN Gambar  | ⏳ (belum) |
| 📄 Koleksi Postman Final         | ⏳ (belum) |
| 🔧 Dashboard Frontend (Admin UI) | ⏳ (belum) |
| 🔒 Mode Maintenance              | ⏳ (opsional) |
| 💬 Sistem Komentar               | ❌ (belum direncanakan) |

---

## 🔌 API Endpoint Highlights

- `POST /auth/register` → Registrasi user (default role: operator)  
- `POST /auth/login` → Autentikasi + JWT  
- `POST /auth/refresh` → Perbarui token  
- `GET /user/me` → Ambil profil dari token  
- `GET /articles`, `GET /articles/:id`, `GET /articles/slug/:slug`  
- `POST /articles` + Gambar (form-data)  
- `GET /categories`, `GET /tags`  
- `GET /pages/:slug` → Halaman statis publik  
- `GET /admin/users`, `PUT /admin/users/:id/role`  
- `GET|POST|PUT|DELETE /admin/roles`

---

## 📦 Cara Menjalankan (Local Development)

```bash
# clone repositori
git clone https://github.com/personalmukti/cms-go.git
cd cms-go

# konfigurasi environment
cp .env.example .env
# atau buat manual sesuai kebutuhan
# contoh isi .env:
# APP_PORT=8080
# DB_HOST=localhost
# DB_PORT=5432
# DB_USER=postgres
# DB_PASSWORD=secretpass
# DB_NAME=cms_go_base
# JWT_SECRET=secret

# install dependensi
go mod tidy

# buat folder uploads untuk simpan gambar
mkdir uploads

# jalankan server
go run main.go
```

---

## 🤝 Kontribusi

Saat ini pengembangan dilakukan oleh pengembang internal. Namun struktur proyek sudah terbuka dan siap dikembangkan lebih jauh (baik untuk keperluan frontend maupun integrasi lanjutan).

---

## 📘 Lisensi

Proyek ini didistribusikan secara terbuka untuk keperluan pribadi, pembelajaran, atau pengembangan mandiri.

---
