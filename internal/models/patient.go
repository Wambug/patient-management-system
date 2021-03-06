package models

// Patient model struct
import "time"

type Patient struct {
	Patientid          int
	Username           string
	Full_name          string
	Email              string
	Dob                time.Time
	Contact            string
	Bloodgroup         string
	Hashed_password    string
	Password_change_at time.Time
	Created_at         time.Time
	//verified           bool
}

//Update Patient strucy
type UpdatePatient struct {
	Id                 int
	Username           string
	Full_name          string
	Email              string
	Dob                time.Time
	Contact            string
	Bloodgroup         string
	Hashed_password    string
	Password_change_at time.Time
}

//PatientRepository represent the Patient repository contract
type PatientRepository interface {
	Create(patient Patient) (Patient, error)
	Find(id int) (Patient, error)
	FindAll() ([]Patient, error)
	Delete(id int) error
	Update(patient UpdatePatient, id int) (Patient, error)
}
