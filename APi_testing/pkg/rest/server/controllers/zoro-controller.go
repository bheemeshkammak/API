package controllers

import (
	"github.com/bheemeshkammak/API/api_testing/pkg/rest/server/models"
	"github.com/bheemeshkammak/API/api_testing/pkg/rest/server/services"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type ZoroController struct {
	zoroService *services.ZoroService
}

func NewZoroController() (*ZoroController, error) {
	zoroService, err := services.NewZoroService()
	if err != nil {
		return nil, err
	}
	return &ZoroController{
		zoroService: zoroService,
	}, nil
}

func (zoroController *ZoroController) CreateZoro(context *gin.Context) {
	// validate input
	var input models.Zoro
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// trigger zoro creation
	if _, err := zoroController.zoroService.CreateZoro(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Zoro created successfully"})
}

func (zoroController *ZoroController) UpdateZoro(context *gin.Context) {
	// validate input
	var input models.Zoro
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger zoro update
	if _, err := zoroController.zoroService.UpdateZoro(id, &input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Zoro updated successfully"})
}

func (zoroController *ZoroController) FetchZoro(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger zoro fetching
	zoro, err := zoroController.zoroService.GetZoro(id)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, zoro)
}

func (zoroController *ZoroController) DeleteZoro(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger zoro deletion
	if err := zoroController.zoroService.DeleteZoro(id); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Zoro deleted successfully",
	})
}

func (zoroController *ZoroController) ListZoros(context *gin.Context) {
	// trigger all zoros fetching
	zoros, err := zoroController.zoroService.ListZoros()
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, zoros)
}

func (*ZoroController) PatchZoro(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "PATCH",
	})
}

func (*ZoroController) OptionsZoro(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "OPTIONS",
	})
}

func (*ZoroController) HeadZoro(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "HEAD",
	})
}
