package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

type Config struct {
	Server struct {
		Port int
	}
	Databases []DBConfig
	JWT       struct {
		Secret string
	}
}

var AppConfig *Config

func main() {
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	AppConfig = &Config{}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Fatalf("Unable to decode into struct: %v", err)
	}

	fmt.Println("Server Port::", AppConfig.Server.Port)
	fmt.Println("First DB name:", AppConfig.Databases[0].Name)
	fmt.Println("Second DB port:", AppConfig.Databases[1].Port)
	fmt.Println("JWT Secret:", AppConfig.JWT.Secret)
	fmt.Println("DB Host:", AppConfig.Databases[0].Host)

	// ✅ Loop through databases
	for i, db := range AppConfig.Databases {
		fmt.Printf("Database #%d:\n", i+1)
		fmt.Println("  Host:", db.Host)
		fmt.Println("  Port:", db.Port)
		fmt.Println("  User:", db.User)
		fmt.Println("  Password:", db.Password)
		fmt.Println("  Name:", db.Name)
	}
}
