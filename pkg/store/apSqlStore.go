package store

import (
	"database/sql"
	"log"

	"github.com/rows15/CTDGolangBackendIII/config"
	"github.com/rows15/CTDGolangBackendIII/internal/domain"
)

func init() {
	config.LoadConfig()
}

// ApStore - Set the contract for ApStore that is made of a composition of Store interface.
type ApStore interface {
	Store
	GetAllAppointmentsByPatientIdentify(identifyNumber string) ([]domain.AppointmentDTO, error)
	GetAllAppointmentsByDentistsLicense(licenseNumber string) ([]domain.AppointmentDTO, error)
	GetAllAppointmentsByDateTimeInterval(startDateTime, endDateTime string) ([]domain.Appointment, error)
}

// NewSQLAp - Initialize ApStore interface
func NewSQLAp() ApStore {
	database, err := config.ConnectDatabase()
	if err != nil {
		panic(err)
	}
	return &appointmentStore{
		sqlStore: &sqlStore{db: database},
		db:       database,
	}
}

type appointmentStore struct {
	*sqlStore
	db *sql.DB
}

// GetAllAppointmentsByPatientIdentify - return a list of all appointments made by a patient through your identity number
func (sa *appointmentStore) GetAllAppointmentsByPatientIdentify(identifyNumber string) ([]domain.AppointmentDTO, error) {
	var appointment domain.AppointmentDTO
	var appointments []domain.AppointmentDTO

	query := "SELECT a.id, a.description, DATE_FORMAT(a.date_and_time,'%d/%m/%Y %H:%i') date_and_time,a.dentist_license,a.patient_identity,d.id,d.surname,d.name,d.license_number,p.id,p.surname,p.name,p.identity_number,DATE_FORMAT(p.created_at,'%d/%m/%Y %H:%i') created_at FROM appointments a INNER JOIN dentists d on a.dentist_license = d.license_number INNER JOIN patients p on a.patient_identity = p.identity_number WHERE a.patient_identity = ? ORDER BY a.date_and_time"
	rows, err := sa.db.Query(query, identifyNumber)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		if err := rows.Scan(
			&appointment.Id,
			&appointment.Description,
			&appointment.DateAndTime,
			&appointment.DentistLicense,
			&appointment.PatientIdentity,
			&appointment.Dentist.Id,
			&appointment.Dentist.Surname,
			&appointment.Dentist.Name,
			&appointment.Dentist.LicenseNumber,
			&appointment.Patient.Id,
			&appointment.Patient.Surname,
			&appointment.Patient.Name,
			&appointment.Patient.IdentityNumber,
			&appointment.Patient.CreatedAt); err != nil {
			return appointments, err
		}
		appointments = append(appointments, appointment)
	}
	return appointments, nil
}

// GetAllAppointmentsByDentistsLicense - return a list of all appointments made by a dentist through your license number
func (sa *appointmentStore) GetAllAppointmentsByDentistsLicense(licenseNumber string) ([]domain.AppointmentDTO, error) {
	var appointment domain.AppointmentDTO
	var appointments []domain.AppointmentDTO

	query := "SELECT a.id, a.description, DATE_FORMAT(a.date_and_time,'%d/%m/%Y %H:%i') date_and_time,a.dentist_license,a.patient_identity,d.id,d.surname,d.name,d.license_number,p.id,p.surname,p.name,p.identity_number,DATE_FORMAT(p.created_at,'%d/%m/%Y %H:%i') created_at FROM appointments a INNER JOIN dentists d on a.dentist_license = d.license_number INNER JOIN patients p on a.patient_identity = p.identity_number WHERE a.dentist_license = ? ORDER BY a.date_and_time"
	rows, err := sa.db.Query(query, licenseNumber)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		if err := rows.Scan(
			&appointment.Id,
			&appointment.Description,
			&appointment.DateAndTime,
			&appointment.DentistLicense,
			&appointment.PatientIdentity,
			&appointment.Dentist.Id,
			&appointment.Dentist.Surname,
			&appointment.Dentist.Name,
			&appointment.Dentist.LicenseNumber,
			&appointment.Patient.Id,
			&appointment.Patient.Surname,
			&appointment.Patient.Name,
			&appointment.Patient.IdentityNumber,
			&appointment.Patient.CreatedAt); err != nil {
			return appointments, err
		}
		appointments = append(appointments, appointment)
	}
	return appointments, nil
}

// GetAllAppointmentsByDateTimeInterval - return a list of all appointments during a datetime interval. Used mostly to validate if a date is available.
func (sa *appointmentStore) GetAllAppointmentsByDateTimeInterval(startDateTime, endDateTime string) ([]domain.Appointment, error) {
	var appointment domain.Appointment
	var appointments []domain.Appointment
	rows, err := sa.db.Query("SELECT * FROM appointments WHERE date_and_time BETWEEN ? AND ?", startDateTime, endDateTime)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		if err := rows.Scan(
			&appointment.Id,
			&appointment.Description,
			&appointment.DateAndTime,
			&appointment.DentistLicense,
			&appointment.PatientIdentity); err != nil {
			return appointments, err
		}
		appointments = append(appointments, appointment)
	}
	return appointments, nil
}
