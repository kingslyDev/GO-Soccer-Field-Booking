package config

import (
	"os"
	"user-services/common/util"

	"github.com/sirupsen/logrus"
)

// Struktur utama untuk menyimpan konfigurasi aplikasi
type AppConfig struct {
	Port                  int            `json:"port"`
	AppName               string         `json:"appName"`
	AppEnv                string         `json:"appEnv"`
	SignatureKey          string         `json:"signatureKey"`
	Database              DatabaseConfig  `json:"database"`
	RateLimiterMaxRequest float64         `json:"rateLimiterMaxRequest"`
	RateLimiterTimeSecond int            `json:"rateLimiterTimeSecond"`
	JwtSecretKey          string         `json:"jwtSecretKey"`
	JwtExpirationTime     int            `json:"jwtExpirationTime"`
}

// Struktur untuk konfigurasi database
type DatabaseConfig struct {
	Host                string `json:"host"`
	Port                int    `json:"port"`
	Name                string `json:"name"`
	Username            string `json:"username"`
	Password            string `json:"password"`
	MaxOpenConnection   int    `json:"maxOpenConnection"`
	MaxLifetimeConnection int  `json:"maxLifetimeConnection"`
	MaxIdleConnection   int    `json:"maxIdleConnection"`
	MaxIdleTime         int    `json:"maxIdleTime"`
}

var config AppConfig

// Fungsi untuk memuat konfigurasi dari file JSON
func init(){
	err := util.BindFromJson(&config, "config", ".") 
	if err != nil {
		logrus.Infof("Failed to load config: %v", err)
		err = util.BindFromJson(&config, os.Getenv("CONSUL_HTTP_URL"), os.Getenv("CONSUL_HTTP_KEY"))
		if err != nil {
			panic(err)
		}
	}
}

// Fungsi untuk mendapatkan konfigurasi
func GetConfig() AppConfig {
	return config
}
