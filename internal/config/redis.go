package config

type Redis struct {
	Addr     string `json:"addr" env:"REDIS_ADDR"`
	Password string `json:"password" env:"REDIS_PASSWORD"`
	DB       int    `json:"db" env:"REDIS_DB"`
}
