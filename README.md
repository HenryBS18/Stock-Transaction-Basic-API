# REST API Sekuritas Saham

REST API sederhana untuk sekuritas saham, dimana user bisa melihat detail saham, melakukan transaksi jual/beli, melihat portfolio dan juga history transaksi. 

### Instruksi Penggunaan:
Disini kita dapat menggunakan XAMPP sebagai server untuk database.

1. Buka XAMPP, lalu nyalakan MySQL.

2. Buka Terminal, lalu buat pindah ke direktori C:\xampp\mysql\bin

```shell
cd C:\xampp\mysql\bin
```

3. Buat database baru dengan nama stock_golang menggunakan perintah:
```sql
CREATE DATABASE stock_golang;
USE stock_golang;
```

4. Buka project REST API, lalu jalankan di terminal:
```go
go run main.go
```
setelah dijalankan, tabel akan otomatis terbuat sesuai pada skema database yang ada pada folder models.

5. Buat data dummy pada database dengan perintah:
```sql
INSERT INTO accounts (name) VALUES ('Henry');
INSERT INTO accounts (name) VALUES ('Andy');

INSERT INTO stocks (ticker, last_price, previous_price, open_price, volume, frequency, turnover)
VALUES
    ('BBCA', 9500, 9500, 9500, 30000000, 2230, 50300000),
    ('BMRI', 4500, 4500, 4500, 6200000, 583, 8543000);
```

6. Aplikasi REST API Sekuritas Saham siap digunakan.
7. Silahkan membuka dokumentasi pada file docs.pdf, untuk petunjuk cara penggunaan endpoint beserta response nya.