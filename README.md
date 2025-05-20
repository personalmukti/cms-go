# CMS-Go

📘 **CMS-Go** adalah sistem manajemen konten (Content Management System) berbasis backend **Go (Golang)** yang dikembangkan sebagai alternatif ringan dan fleksibel dari platform seperti WordPress atau Blogspot. Proyek ini bertujuan membangun CMS modular, mudah dikembangkan, dan dapat digunakan secara mandiri tanpa layanan pihak ketiga berbayar.

---

## 🚀 Status Pengembangan

> Proyek ini **masih dalam tahap awal pengembangan**. Struktur dasar sedang dibangun, termasuk fitur-fitur inti seperti manajemen artikel, kategori, dan pengguna.

---

## 🔧 Teknologi yang Digunakan

| Teknologi      | Deskripsi                                      |
|----------------|-----------------------------------------------|
| [Go](https://go.dev/)             | Bahasa pemrograman utama                  |
| [Echo](https://echo.labstack.com/)           | Web framework untuk routing & middleware |
| [PostgreSQL](https://www.postgresql.org/)     | Database utama (bisa diganti jika perlu) |
| JWT (JSON Web Token)              | Autentikasi & otorisasi user             |
| GORM                              | ORM untuk pengelolaan database           |
| Docker (opsional)                | Untuk pengembangan berbasis container    |

---

## 📂 Struktur Modul (Rencana Awal)

- `auth/` → Modul autentikasi & user management
- `posts/` → Artikel, editor, slug, thumbnail
- `categories/` → Manajemen kategori & tag
- `comments/` → Sistem komentar (jika diaktifkan)
- `media/` → Upload dan pengelolaan file
- `pages/` → Halaman statis seperti Tentang/Kontak
- `settings/` → Konfigurasi situs & metadata
- `admin/` → Dashboard backend

---

## ⚙️ Fitur Utama (Roadmap)

- [ ] Manajemen Artikel (CRUD)
- [ ] Manajemen Kategori & Tag
- [ ] Upload Gambar & Media
- [ ] Sistem Autentikasi (JWT)
- [ ] Halaman Statis
- [ ] Dashboard Admin
- [ ] Role-based User Access
- [ ] API Publik dan Terproteksi
- [ ] Mode Maintenance

---

## 📦 Cara Menjalankan (Development)

```bash
# clone repositori
git clone https://github.com/personalmukti/cms-go.git
cd cms-go

# konfigurasi .env
cp .env.example .env

# install dependencies (Go modules)
go mod tidy

# jalankan server
go run main.go
