package customer

import "time"

type CustomerData struct {
	FirstName   string
	LastName    string
	DateofBirth string
	Married     bool
	PhoneNumber string
	EmailAddr   string
	CreatedDate time.Time
	Tags        []string
}
