package infra

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kvngho/clayful-monitoring/internal/domain/ent"
	"github.com/spf13/viper"
)

func EntConfigFromENV() EntConfig {
	return EntConfig{
		Host: viper.GetString("DATABASE_HOST"),
		Port: viper.GetInt("DATABASE_PORT"),
		User: viper.GetString("DATABASE_USER"),
		Password: viper.GetString("DATABASE_PASSWORD"),
		DBName: viper.GetString("DATABASE_NAME"),
	}
	// var cfg EntConfig

	// if err := env.Parse(&cfg); err != nil {
	// 	panic(err)
	// }
	// fmt.Println(cfg)
	// return cfg
}

type EntConfig struct {
	Host 		string `env:"DATABASE_HOST"`
	Port 		int	   `env:"DATABASE_PORT"`
	User 		string `env:"DATABASE_USER"`
	Password 	string `env:"DATABASE_PASSWORD"`
	DBName 		string `env:"DATABASE_NAME"`
}

type EntDB struct {
	DB *ent.Client
}

func NewEntCilent(cfg EntConfig) EntDB {
	client, err := ent.Open("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=True", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName))
	if err != nil {
		log.Fatalln("database connection error")
	}
	return EntDB{
		DB: client,
	}
}