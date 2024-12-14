package app

import (
	"flag"
	"log"

	"github.com/frangklynbfndruru/e-commerce/app/configuration"
	"github.com/joho/godotenv"

	"fmt"
	"os"
)

func RunApp() {
	var server = configuration.Server{}
	var appConfig = configuration.AppConfig{}
	var dbConfig = configuration.DbConfig{}

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
		fmt.Println("error .env file!")
	}

	appConfig.AppName = os.Getenv("APP_NAME")
	appConfig.AppEnv = os.Getenv("APP_ENV")
	appConfig.AppPort = os.Getenv("APP_PORT")

	dbConfig.DbHost = os.Getenv("DB_HOST")
	dbConfig.DbUser = os.Getenv("DB_USER")
	dbConfig.DbPassword = os.Getenv("DB_PASSWORD")
	dbConfig.DbName = os.Getenv("DB_NAME")
	dbConfig.DbPort = os.Getenv("DB_PORT")

	flag.Parse() //untuk menerima command go run dari terminal
	// arg := flag.Arg(0) //mengambil argumen pertama dari command line. contoh `go run db:migrate`

	// if arg != "" {
	// 	server.InitCommands(appConfig, dbConfig)

	// } else {
	server.Initialize(appConfig, dbConfig)
	server.RunDefaultPort(":" + appConfig.AppPort)
	// }
}
