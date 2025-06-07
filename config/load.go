package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

func LoadConfig() {
	viper.SetConfigType("env")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	viper.SetDefault("MONGO_URL", "mongodb://mongo:icBbXHqaqUoyMPYZijCkurGaLzPZsvKB@mongodb.railway.internal:27017")

	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Println("⚠️  .env não encontrado, usando variáveis do sistema")
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Fatalf("❌ Erro ao carregar config: %v", err)
	}

	log.Println("✅ Config carregada com sucesso")
}
