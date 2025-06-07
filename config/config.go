package config

type Config struct {
	MongoURL string `mapstructure:"MONGO_URL"`
	Port     string `mapstructure:"PORT"`
}

var AppConfig Config
