# 📦 Microservices Project with Kafka & PostgreSQL

Proyek ini menggunakan arsitektur microservices dengan beberapa layanan yang berjalan secara terpisah dan berkomunikasi melalui Kafka. Database yang digunakan adalah PostgreSQL.

## 🚀 Cara Menjalankan

### 1️⃣ Persyaratan
Pastikan Anda sudah menginstal:
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

### 2️⃣ Clone Repository
```sh
git clone https://github.com/username/repo.git
cd repo
```

### 3️⃣ Buat File `.env`
Buat file `.env` di root proyek dan isi dengan variabel lingkungan yang diperlukan.

Contoh:
```env
DOCKER_HOST_IP=127.0.0.1
```

### 4️⃣ Jalankan Docker Compose
Jalankan perintah berikut untuk membangun dan menjalankan layanan:
```sh
docker-compose up -d
```

📌 **Opsi tambahan**:
- Gunakan `docker-compose up --build -d` jika ingin membangun ulang image.
- Gunakan `docker-compose down` untuk menghentikan semua layanan.

### 5️⃣ Cek Layanan yang Berjalan
```sh
docker ps
```
Pastikan semua container berjalan tanpa error.

### 6️⃣ Mengakses Layanan

| Layanan       | URL |
|--------------|--------------------------------|
| Kafka UI     | [http://localhost:8070](http://localhost:8070) |
| PgAdmin      | [http://localhost:5050](http://localhost:5050) |

### 7️⃣ Mengelola PostgreSQL
Untuk mengakses database menggunakan PgAdmin:
1. Masuk ke `http://localhost:5050`
2. Login dengan:
   - **Email**: `admin@pgadmin.com`
   - **Password**: `admin`
3. Tambahkan Server:
   - **Host**: `postgres`
   - **Username**: `root`
   - **Password**: `barca1899`

### 8️⃣ Logs & Debugging
- **Melihat logs container**:
  ```sh
  docker-compose logs -f
  ```
- **Masuk ke container PostgreSQL**:
  ```sh
  docker exec -it postgres psql -U root
  ```

---

Jika ada error atau masalah, pastikan semua dependensi telah diinstal dengan benar dan coba jalankan ulang dengan:

```sh
docker-compose down -v
docker-compose up --build -d
```

🎯 **Happy Explore My Project!** 🚀

