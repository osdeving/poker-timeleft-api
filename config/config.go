package config

type Config struct {
	MongoURI string `mapstructure:"MONGO_PUBLIC_URL"`
}

var AppConfig Config
