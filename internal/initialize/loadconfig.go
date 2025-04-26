package initialize

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"github.com/thinhcompany/go-ecommerce-backend-api/global"
)

func LoadConfig() {
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	if err := viper.Unmarshal(&global.Config); err != nil {
		log.Fatalf("Unable to decode into struct: %v", err)
	}

	fmt.Println("Server running on port:", global.Config.Server.Port)
	fmt.Println("Database name:", global.Config.MySQL.DBName)
	fmt.Println("JWT secret:", global.Config.JWT.Secret)

}
