package handlers

import (
	"net/http"
	"strconv"

	"github.com/0xDevvvvv/makerble/internal/models"
	"github.com/0xDevvvvv/makerble/internal/repositories"
	"github.com/gin-gonic/gin"
)

type PatientHandler struct {
	repo repositories.PatientRepository
}

func NewPatientHandler(repo repositories.PatientRepository) *PatientHandler {
	return &PatientHandler{repo: repo}
}

// CreatePatient handles POST /patients
func (h *PatientHandler) CreatePatient(c *gin.Context) {
	var patient models.Patient
	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid patient data"})
		return
	}

	createdPatient, err := h.repo.Create(&patient)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create patient"})
		return
	}

	c.JSON(http.StatusCreated, createdPatient)
}

// GetPatient handles GET /patients/:id
func (h *PatientHandler) GetPatient(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid patient ID"})
		return
	}

	patient, err := h.repo.GetById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve patient"})
		return
	}
	if patient == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	c.JSON(http.StatusOK, patient)
}

// UpdatePatient handles PUT /patients
func (h *PatientHandler) UpdatePatient(c *gin.Context) {
	var patient models.Patient
	if err := c.ShouldBindJSON(&patient); err != nil || patient.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid patient data"})
		return
	}

	err := h.repo.Update(&patient)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update patient"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Patient updated successfully"})
}

// DeletePatient handles DELETE /patients
func (h *PatientHandler) DeletePatient(c *gin.Context) {
	var patient models.Patient
	if err := c.ShouldBindJSON(&patient); err != nil || patient.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid patient ID"})
		return
	}

	err := h.repo.Delete(&patient)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete patient"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Patient deleted successfully"})
}

// GetAllPatient handles GET /patients
func (h *PatientHandler) GetAllPatient(c *gin.Context) {
	patients, err := h.repo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch patients"})
		return
	}

	c.JSON(http.StatusOK, patients)
}
