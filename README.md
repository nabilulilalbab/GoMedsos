# GoMedsos

GoMedsos adalah proyek backend berbasis Go yang dirancang untuk mendukung sistem media sosial. Proyek ini terdiri dari beberapa service terpisah untuk memudahkan skalabilitas dan pemeliharaan, seperti `imageService`, `userService`, dan `library`.

---

## 📁 Struktur Direktori

```bash
GoMedsos/
│
├── imageService/     # Layanan untuk manajemen gambar
│   ├── main.go
│   ├── go.mod
│   └── go.sum
│
├── userService/      # (Placeholder) Layanan untuk manajemen user
│   └── ...
│
├── library/          # Modul pustaka bersama (helper, response, dll)
│   ├── go.mod
│   ├── helper.go
│   └── response.go

