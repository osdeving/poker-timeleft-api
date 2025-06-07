package config

type Config struct {
	MongoURL string `mapstructure:"MONGO_URL"`
}

var AppConfig Config
