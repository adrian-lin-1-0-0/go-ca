package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/adrian-lin-1-0-0/go-ca/domain"
	swagger "github.com/adrian-lin-1-0-0/go-ca/swagger-codegen/go"
)

type PetHandler struct {
	petUsecase domain.PetUsecase
}

func NewPetHandler(opts *PetHandlerOptions) *PetHandler {
	return &PetHandler{
		petUsecase: opts.Usecase,
	}
}

type PetHandlerOptions struct {
	Usecase domain.PetUsecase
}

type RegisterPetHandlerOptions struct {
	Engine  *gin.Engine
	Handler *PetHandler
}

func RegisterPetHandler(opts *RegisterPetHandlerOptions) {

	opts.Engine.POST("/v2/pet", opts.Handler.AddPet)
	opts.Engine.PUT("/v2/pet", opts.Handler.UpdatePet)
	opts.Engine.GET("/v2/pet/:petId", opts.Handler.GetPetById)
	opts.Engine.DELETE("/v2/pet/:petId", opts.Handler.DeletePet)
}

func (h *PetHandler) GetPetById(c *gin.Context) {
	petId := c.Param("petId")

	pet, err := h.petUsecase.GetPetById(c, petId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &swagger.Pet{
		Id:   pet.Id,
		Name: pet.Name,
	})
}

func (h *PetHandler) AddPet(c *gin.Context) {
	var pet swagger.Pet
	if err := c.ShouldBindJSON(&pet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := h.petUsecase.AddPet(c, &domain.Pet{
		Id:   pet.Id,
		Name: pet.Name,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "AddPet"})
}

func (h *PetHandler) DeletePet(c *gin.Context) {
	petId := c.Param("petId")

	err := h.petUsecase.DeletePet(c, petId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "DeletePet"})
}

func (h *PetHandler) UpdatePet(c *gin.Context) {
	var pet swagger.Pet
	if err := c.ShouldBindJSON(&pet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := h.petUsecase.UpdatePet(c, &domain.Pet{
		Id:   pet.Id,
		Name: pet.Name,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "UpdatePet"})
}
