package dentist

import (
	"errors"
	"log"

	"github.com/rows15/CTDGolangBackendIII/internal/domain"
	"github.com/rows15/CTDGolangBackendIII/pkg/store"
)

var table = "dentists"

type Repository interface {
	GetAll() (interface{}, error)
	GetByID(id int) (interface{}, error)

	Create(d domain.Dentist) (interface{}, error)
	Update(id int, d domain.Dentist) (interface{}, error)
	Delete(id int) error
}

type repository struct {
	store store.Store
}

func NewRepository(store store.Store) Repository {
	return &repository{store}
}

// GetAll - returns all dentists from database
func (r *repository) GetAll() (interface{}, error) {
	return r.store.GetAll(table)
}

func (r *repository) GetByID(id int) (interface{}, error) {
	return r.store.GetByID(id, table)
}

func (r *repository) Create(d domain.Dentist) (interface{}, error) {
	if !r.validateLicenseNumber(d.LicenseNumber) {
		return nil, errors.New("license number already exists on database")
	}
	return r.store.Save(d, table)
}

func (r *repository) Update(id int, d domain.Dentist) (interface{}, error) {
	var dentists []domain.Dentist
	dentistsInterface, err := r.GetAll()
	if err != nil {
		log.Fatalln("error while trying to fetch data from db")
		return nil, err
	}
	dentists, ok := dentistsInterface.([]domain.Dentist)
	if !ok {

		return nil, err
	}

	for _, dentist := range dentists {
		if dentist.Id == id {

			if !r.validateLicenseNumber(d.LicenseNumber) && d.LicenseNumber != dentist.LicenseNumber {
				return nil, errors.New("license number already exists")
			}

			return r.store.Update(id, d, table)
		}
	}
	return nil, errors.New("dentist not found")
}

func (r *repository) Delete(id int) error {

	return r.store.Delete(id, table)
}

func (r *repository) validateLicenseNumber(licenseNumber string) bool {

	var dentists []domain.Dentist
	dentistsInterface, err := r.GetAll()
	if err != nil {
		return false
	}
	dentists, ok := dentistsInterface.([]domain.Dentist)
	if !ok {
		return false
	}

	for _, dentist := range dentists {

		if dentist.LicenseNumber == licenseNumber {
			return false
		}
	}
	return true
}
