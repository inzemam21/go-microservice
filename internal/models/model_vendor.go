package models

type Vendor struct {
	VendorID string `gorm:"primaryKey" json:"vendor_id"`
	Name     string `json:"name"`
	Contact  string `json:"contact"`
	Phone    string `json:"phone_number"`
	Email    string `json:"email_address"`
	Address  string `json:"address"`
}
