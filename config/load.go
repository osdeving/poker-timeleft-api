package config

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

func LoadConfig() {
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	viper.SetDefault("MONGO_URL", "mongodb://localhost:27017")
	viper.SetDefault("PORT", "8080")

	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	if err := viper.ReadInConfig(); err != nil {
		log.Println("⚠️  .env não encontrado, usando variáveis do sistema")
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Fatalf("❌ Erro ao carregar config: %v", err)
	}

	log.Println("✅ Config carregada com sucesso")

	log.Printf("🔎 OS Env PORT: %s", os.Getenv("PORT"))
	log.Printf("🔎 OS Env MONGO_URL: %s", os.Getenv("MONGO_URL"))

	log.Printf("✅ Viper PORT: %s", AppConfig.Port)
	log.Printf("✅ Viper MONGO_URL: %s", AppConfig.MongoURL)
}
