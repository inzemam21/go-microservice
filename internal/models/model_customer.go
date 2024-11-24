package models

type Customer struct {
	CustomerID string `gorm:"primaryKey" json:"customerId"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email_address"`
	Phone      string `json:"phone_number"`
	Address    string `json:"address"`
}
