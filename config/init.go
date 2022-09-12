package config

type Config struct {
	Port       string `yaml:"port"`
	DBHost     string `yaml:"db.redis_host"`
	DBPassword string `yaml:"db.redis_pass"`
	DBPort     int    `yaml:"db.redis_port"`
	DBUser     string `yaml:"db.redis_db"`
}
