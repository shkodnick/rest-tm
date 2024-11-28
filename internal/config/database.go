package config

type Database struct {
	Host     string `json:"host" env:"POSTGRES_HOST"`
	Port     int    `json:"port" env:"POSTGRES_PORT"`
	Database string `json:"database" env:"POSTGRES_DBNAME"`
	User     string `json:"user" env:"POSTGRES_USER"`
	Pass     string `json:"pass" env:"POSTGRES_PASSWORD"`
}
