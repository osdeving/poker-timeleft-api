package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

func LoadConfig() {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		log.Println("⚠️  .env não encontrado, usando variáveis do sistema")
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Fatalf("❌ Erro ao carregar config: %v", err)
	}

	log.Println("✅ Config carregada com sucesso")
}
