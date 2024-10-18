package usecase

import (
	"github.com/matthewyu246/back/models"
	"github.com/matthewyu246/back/repository"
)

type IPhotoUsecase interface {
	GetPhotoById(id int64) (*models.Photo, error)
	UploadPhoto(photo *models.Photo) error
}

type photoUsecase struct {
	pr repository.IPhotoRepository
}

func NewPhotoUsecase(pr repository.IPhotoRepository) IPhotoUsecase {
	return &photoUsecase{pr}
}

func (pu *photoUsecase) GetPhotoById(id int64) (*models.Photo, error) {
	return pu.pr.GetPhotoById(id)
}

func (pu *photoUsecase) UploadPhoto(photo *models.Photo) error {
	return pu.pr.SavePhoto(photo)
}
