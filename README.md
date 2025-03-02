# REST-API

# Provinces API

Proyek ini adalah aplikasi API sederhana menggunakan Golang dengan framework Gin dan GORM untuk mengelola data provinsi di Indonesia. API ini mengambil data dari sumber eksternal dan menyimpannya ke database MySQL.

## Fitur
- Mengambil data provinsi dari API eksternal
- Menyimpan data ke dalam database MySQL
- Menyediakan endpoint untuk mendapatkan daftar provinsi
- Logging dan error handling

## Instalasi
### Prasyarat
Sebelum menjalankan proyek ini, pastikan Anda sudah menginstal:
- Golang
- MySQL
- Postman (opsional, untuk menguji API)

### Cara Menjalankan
1. Clone repository ini:
   ```sh
   git clone https://github.com/username/provinces-api.git
   cd provinces-api
   ```
2. Instal dependensi:
   ```sh
   go mod tidy
   ```
3. Konfigurasi database di `.env`:
   ```sh
   DB_USER=root
   DB_PASSWORD=password
   DB_NAME=provinces_db
   DB_HOST=localhost
   DB_PORT=3306
   ```
4. Jalankan aplikasi:
   ```sh
   go run main.go
   ```

## Endpoint API
### 1. Ambil daftar provinsi
**Request:**
```
GET /provinces
```
**Response:**
```json
{
  "status": "success",
  "code": 200,
  "message": "Berhasil mendapatkan data",
  "data": [
    {"id": 1, "name": "ACEH"},
    {"id": 2, "name": "SUMATERA UTARA"}
  ]
}
```

### 2. Sinkronisasi data dari API eksternal
**Request:**
```
GET /sync-provinces
```
**Response:**
```json
{
  "status": "success",
  "code": 200,
  "message": "Data berhasil disinkronisasi"
}
```

## Troubleshooting
### 1. Port 8080 Sudah Digunakan
**Error:**
```
[GIN-debug] [ERROR] listen tcp :8080: bind: Only one usage of each socket address
```
**Solusi:**
Matikan proses yang menggunakan port 8080:
```sh
netstat -ano | findstr :8080
```
Hentikan proses dengan ID yang ditemukan:
```sh
taskkill /PID <PID> /F
```
Atau jalankan di port lain dengan mengedit `main.go`:
```go
r.Run(":8081")
```

### 2. API Eksternal Tidak Dapat Dihubungi
**Error:**
```
Failed to fetch API: Get "https://api.example.com/provinces": dial tcp: lookup api.example.com: no such host
```
**Solusi:**
- Pastikan API eksternal aktif dan URL sudah benar.
- Coba akses API eksternal melalui browser untuk memastikan tidak ada masalah koneksi.
- Periksa apakah ada firewall atau VPN yang menghalangi akses.

## Lisensi
Proyek ini menggunakan lisensi MIT. Anda bebas menggunakannya dengan ketentuan yang berlaku.

---
Jika ada pertanyaan atau masalah, silakan ajukan di bagian **Issues** pada repository ini. 😊

