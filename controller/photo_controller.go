package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/matthewyu246/back/models"
	"github.com/matthewyu246/back/usecase"
)

type IPhotoController interface {
	UploadPhoto(c *gin.Context)
	GetPhoto(c *gin.Context)
}

type photoController struct {
	pu usecase.IPhotoUsecase
}

func NewPhotoController(pu usecase.IPhotoUsecase) IPhotoController {
	return &photoController{pu}
}

func (pc *photoController) UploadPhoto(c *gin.Context) {
	file, err := c.FormFile("photo")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}

	f, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer f.Close()

	data := make([]byte, file.Size)
	_, err = f.Read(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	photo := &models.Photo{
		Title: file.Filename,
		Data:  data,
	}

	if err := pc.pu.UploadPhoto(photo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Photo uploaded successfully", "photo_id": photo.ID})
}

func (pc *photoController) GetPhoto(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid photo ID"})
		return
	}

	photo, err := pc.pu.GetPhotoById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
		return
	}

	c.Data(http.StatusOK, "image/jpeg", photo.Data)
}
