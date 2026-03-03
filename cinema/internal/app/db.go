package app

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Dbname   string `mapstructure:"dbname"`
}

func InitDbConnection(databaseConfig Database) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", databaseConfig.Host, databaseConfig.Username, databaseConfig.Password, databaseConfig.Dbname, databaseConfig.Port)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
