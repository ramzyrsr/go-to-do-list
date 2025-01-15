package models

type Product struct {
	ID          int     `json:"id"`
	ApplicantId string  `json:"applicant_id"`
	FirstName   string  `json:"first_name"`
	LastName    string  `json:"last_name"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
}
