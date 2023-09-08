package main

import (
	"github.com/adrian-lin-1-0-0/go-ca/service/delivery/http/v2"
	"github.com/adrian-lin-1-0-0/go-ca/service/usecase"
	"github.com/gin-gonic/gin"
)

func main() {

	engine := gin.Default()

	petUsecase := usecase.NewPetUsecase(
		&usecase.PetUsecaseOptions{
			PetRepo: nil,
		},
	)

	http.RegisterPetHandler(&http.PetHandlerOptions{
		Engine:  engine,
		Usecase: petUsecase,
	})

	engine.Run(":8080")
}
