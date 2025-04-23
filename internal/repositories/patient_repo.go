package repositories

import "github.com/0xDevvvvv/makerble/internal/models"

type PatientRepository interface {
	GetById(id int) (*models.Patient, error)
	Create(patient *models.Patient) error
	Update(patient *models.Patient) error
	Delete(patient *models.Patient) error
}
