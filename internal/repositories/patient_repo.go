package repositories

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/0xDevvvvv/makerble/internal/models"
)

type PatientRepository interface {
	GetById(id int) (*models.Patient, error)
	GetByName(name string) (*models.Patient, error)
	Create(patient *models.Patient) (*models.Patient, error)
	Update(patient *models.Patient) error
	Delete(patient *models.Patient) error
	GetAll() ([]models.Patient, error)
}

type patientRepo struct {
	db *sql.DB
}

func NewPatientRepository(db *sql.DB) PatientRepository {
	return &patientRepo{db}
}

func (r *patientRepo) GetById(id int) (*models.Patient, error) {
	var patient models.Patient
	query := `
		SELECT id, name, age, gender, address, phone, illness, created_at
		FROM patients
		WHERE id = $1;
	`

	row := r.db.QueryRow(query, id)

	err := row.Scan(&patient.ID, &patient.Name, &patient.Age, &patient.Gender, &patient.Address, &patient.Phone, &patient.Illness, &patient.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // patient not found
		}
		return nil, fmt.Errorf("error Fetching Patient (By ID): %w", err)
	}

	return &patient, nil
}

func (r *patientRepo) GetByName(name string) (*models.Patient, error) {
	var patient models.Patient
	query := `
		SELECT id, name, age, gender, address, phone, illness, created_at
		FROM patients
		WHERE name = $1;
	`

	row := r.db.QueryRow(query, name)

	err := row.Scan(&patient.ID, &patient.Name, &patient.Age, &patient.Gender, &patient.Address, &patient.Phone, &patient.Illness, &patient.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // patient not found
		}
		return nil, fmt.Errorf("error Fetching Patient (By Name): %w", err)
	}

	return &patient, nil
}

func (r *patientRepo) Create(patient *models.Patient) (*models.Patient, error) {
	query := `
		INSERT INTO patients (name, age, gender, address, phone, illness, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, created_at;
	`

	err := r.db.QueryRow(query, patient.Name, patient.Age, patient.Gender, patient.Address, patient.Phone, patient.Illness, time.Now()).
		Scan(&patient.ID, &patient.CreatedAt)

	if err != nil {
		return nil, fmt.Errorf("error creating patient: %w", err)
	}

	return patient, nil
}

func (r *patientRepo) Update(patient *models.Patient) error {
	query := `
		UPDATE patients
		SET name = $1, age = $2, gender = $3, address = $4, phone = $5, illness = $6
		WHERE id = $7;
	`

	_, err := r.db.Exec(query, patient.Name, patient.Age, patient.Gender, patient.Address, patient.Phone, patient.Illness, patient.ID)
	if err != nil {
		return fmt.Errorf("error updating patient: %w", err)
	}

	return nil
}

func (r *patientRepo) Delete(patient *models.Patient) error {
	query := `
		DELETE FROM patients
		WHERE id = $1;
	`

	_, err := r.db.Exec(query, patient.ID)
	if err != nil {
		return fmt.Errorf("error deleting patient: %w", err)
	}

	return nil
}
func (r *patientRepo) GetAll() ([]models.Patient, error) {
	query := `
		SELECT id, name, age, gender, address, phone, illness, created_at
		FROM patients;
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error fetching all patients: %w", err)
	}
	defer rows.Close()

	var patients []models.Patient

	for rows.Next() {
		var p models.Patient
		err := rows.Scan(
			&p.ID, &p.Name, &p.Age, &p.Gender,
			&p.Address, &p.Phone, &p.Illness, &p.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning patient row: %w", err)
		}
		patients = append(patients, p)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error in getting all patients: %w", err)
	}

	return patients, nil
}
