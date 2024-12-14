package configuration

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/urfave/cli"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

/*
Server bertanggung jawab untuk mengatur kompnen dalam menjalankan aplikasi
-Koneksi kedalam DB
-Routing HTTP
Konfigurasi Aplikasi
*/
type Server struct {
	//variabel & library yang digunakan
	DB        *gorm.DB //ORM menggunakan GORM
	Router    *mux.Router
	AppConfig *AppConfig
}

/*
AppConfig bertujuan untuk insialisasi/menyimpan konfigurasi aplikasi sehingga
lebih practice saat dijalankan dan akan disimpan kedalam Server
*/
type AppConfig struct {
	AppName string
	AppEnv  string
	AppPort string
	AppURl  string
}

/*
DbConfig bertujuan untuk insialisasi/menyimpan konfigurasi databse sehingga
lebih practice saat dijalankan dan akan disimpan kedalam Server
*/
type DbConfig struct {
	DbHost     string
	DbUser     string
	DbPassword string
	DbName     string
	DbPort     string
}

func (server *Server) RunDefaultPort(address string) {

	fmt.Printf("Listening to port %s", address)

	log.Fatal(http.ListenAndServe(address, server.Router))
}

// function ketika aplikasi sudah jalan
func (server *Server) Initialize(appConfig AppConfig, dbConfig DbConfig) {
	fmt.Println("Welcome to " + appConfig.AppName)
}

func (server *Server) InitializeDB(dbConfig DbConfig) {
	var err error

	//dsn = Data Source Name
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/jakarta",
		dbConfig.DbHost, dbConfig.DbUser, dbConfig.DbPassword, dbConfig.DbName, dbConfig.DbPort)

	server.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
}

func (server *Server) dbMigrate() {

}


/*
Fungsi initCommands bertugas untuk:

	1. Menghubungkan ke database menggunakan pengaturan yang diberikan.
	2. Menyiapkan aplikasi Command-Line (CLI) yang bisa menerima perintah dari terminal.
	3. Menangani kesalahan: jika ada masalah saat menjalankan perintah CLI, program akan menampilkan pesan error dan berhenti.

Jadi, fungsi ini membantu server menjalankan perintah melalui terminal sambil memastikan database sudah siap.
*/
func (server *Server) initCommands(appConfig AppConfig, dbConfig DbConfig) {
	server.InitializeDB(dbConfig)//=>1
	commandApp := cli.NewApp()//=>2
	err := commandApp.Run(os.Args)//=>3
	if err != nil {
		log.Fatal(err)
	}
}
