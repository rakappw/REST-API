# REST-API

# Provinces API

Proyek ini adalah aplikasi API sederhana menggunakan Golang dengan framework Gin dan GORM untuk mengelola data provinsi di Indonesia. API ini mengambil data dari sumber eksternal dan menyimpannya ke database MySQL.

## Fitur
- Mengambil data provinsi dari API eksternal
- Menyimpan data ke dalam database MySQL
- Menyediakan endpoint untuk mendapatkan daftar provinsi

## Instalasi
### Prasyarat
Sebelum menjalankan proyek ini, pastikan Anda sudah menginstal:
- Golang
- MySQL
- Postman (opsional, untuk menguji API)

### Cara Menjalankan
1. Clone repository ini:
   ```sh
   git clone https://github.com/rakappw/provinces-api.git
   cd provinces-api
   ```
2. Instal dependensi:
   ```sh
   go mod tidy
   ```
3. Jalankan aplikasi:
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


