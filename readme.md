# QuickCare

One Stop Solution untuk Akses Kesehatan Lebih Mudah dengan Rujukan Online dan Pengelolaan Obat yang Efisien

## Product Description

QuickCare adalah solusi kesehatan digital yang mempermudah proses rujukan online dan manajemen obat secara terintegrasi, sehingga pasien dan keluarga dapat mengakses layanan kesehatan dengan cepat dan efektif.

## Features

### 1. Rujukan Online

**Fitur Rujukan Online memungkinkan pasien untuk mengajukan dan memperpanjang rujukan secara digital**

Fitur ini memungkinkan pasien untuk melakukan rujukan secara online dengan mengirimkan video dan keterangan langsung kepada puskesmas, yang dapat berkoordinasi dengan rumah sakit. Pasien juga dapat memperpanjang rujukan secara otomatis melalui integrasi dengan rekam medis, sehingga proses rujukan menjadi lebih cepat dan efisien.

### 2. Manajemen Obat

**Fitur manajemen obat untuk mengelola informasi obat dan memudahkan pengingat konsumsi obat secara digital**

Fitur ini terintegrasi dengan rekam medis dan memungkinkan keluarga untuk memantau jadwal konsumsi obat pasien. Sistem akan mencatat inventaris obat, memberi pengingat otomatis sesuai jadwal, dan memungkinkan pasien untuk mengaktifkan atau menonaktifkan obat yang dikonsumsi, membantu memastikan kepatuhan dan kontrol yang lebih baik terhadap pengobatan.

## Team PIJAR

* Michael Sihotang - Institut Teknologi Bandung
* Erensi Ratu Chelsia - Institut Teknologi Bandung
* Abraham Megantoro Samudra - Institut Teknologi Bandung

## Getting Started

### Installation & Running the App

1. Clone repository
```bash
git clone https://github.com/abrahammegantoro/quickcare-bpjs-be
cd quickcare-bpjs-be
```

2. Install dependencies
```bash
go mod vendor
```

3. Start the development server
```bash
go run cmd/api/main.go
```

## Project Structure

```
.
├── Dockerfile
├── cmd
│   └── api
│       └── main.go
├── config
│   └── config.go
├── db
│   └── database.go
├── go.mod
├── go.sum
├── handlers
│   ├── medicine.go
│   └── referral.go
├── middlewares
├── models
│   ├── auth.go
│   ├── medicine.go
│   ├── model.go
│   ├── referral.go
│   └── user.go
├── repositories
│   ├── medicine.go
│   ├── referral.go
│   └── user.go
├── usecases
│   ├── medicine.go
│   ├── referral.go
│   └── user.go
└── utils
```

## Deployment

[**Android**](https://expo.dev/accounts/abrahamsamudra/projects/quickcare-bpjs-fe/builds/5f34a6ab-dfad-4871-aca2-ad9af7755177)

## Related Repositories

[**Frontend Repository**](https://github.com/abrahammegantoro/quickcare-bpjs-fe)