package config

type Config struct {
	MongoURI string `mapstructure:"MONGO_URI"`
}

var AppConfig Config
