package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nguyendangviet12022004/cinema-microservices/cinema/internal/app"
	"github.com/spf13/viper"
)

func main() {

	// init viper
	app.InitConfig()

	// get db config
	var dbConfig app.Database
	err := viper.UnmarshalKey("database", &dbConfig)
	if err != nil {
		panic(err)
	}

	// init db connection
	db, err := app.InitDbConnection(dbConfig)
	if err != nil {
		panic(err)
	}

	// database migration
	err = db.AutoMigrate()
	if err != nil {
		panic(err)
	}

	//run server
	r := gin.Default()

	r.Run(":" + viper.GetString("server.port"))
}
