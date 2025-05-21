package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("Gagal memuat .env file (default gunakan variable kosong.)")
	}
}

func Get(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Printf("⚠️  Peringatan: %s tidak ditemukan di environment", key)
	}
	return value
}