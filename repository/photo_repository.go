package repository

import (
	"github.com/matthewyu246/back/models"
	"gorm.io/gorm"
)

type IPhotoRepository interface {
	SavePhoto(photo *models.Photo) error
	GetPhotoById(id int64) (*models.Photo, error)
}

type photoRepository struct {
	db *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) IPhotoRepository {
	return &photoRepository{db}
}

func (pr *photoRepository) SavePhoto(photo *models.Photo) error {
	if err := pr.db.Create(photo).Error; err != nil {
		return err
	}
	return nil
}

func (pr *photoRepository) GetPhotoById(id int64) (*models.Photo, error) {
	var photo models.Photo
	if err := pr.db.First(&photo, id).Error; err != nil {
		return nil, err
	}
	return &photo, nil
}
