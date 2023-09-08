package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/adrian-lin-1-0-0/go-ca/service/delivery/http/v2"
	"github.com/adrian-lin-1-0-0/go-ca/service/repository/mysql"
	"github.com/adrian-lin-1-0-0/go-ca/service/usecase"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	viper.SetConfigFile(".env")
	viper.SetConfigType("dotenv")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Fatal error config file: %v\n", err)
	}
}

func main() {
	dbHost := viper.GetString("DB_HOST")
	dbDatabase := viper.GetString("DB_DATABASE")
	dbUser := viper.GetString("DB_USER")
	dbPwd := viper.GetString("DB_PASSWORD")
	dbPort := viper.GetString("DB_PORT")

	db, err := sql.Open(
		"mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPwd, dbHost, dbPort, dbDatabase),
	)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	petRepo := mysql.NewPetRepository(db)

	petUsecase := usecase.NewPetUsecase(
		&usecase.PetUsecaseOptions{
			PetRepo: petRepo,
		},
	)

	engine := gin.Default()

	http.RegisterPetHandler(&http.PetHandlerOptions{
		Engine:  engine,
		Usecase: petUsecase,
	})

	engine.Run(":8080")
}
