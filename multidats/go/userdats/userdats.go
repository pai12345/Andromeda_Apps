package userdats

import (
	"fmt"
	"time"
)

type userData struct {
	firstName   string
	lastName    string
	dateofBirth string
	married     bool
	phoneNumber string
	createdDate time.Time
}

func GenerateUserdata(firstname string, lastname string, dob string, married bool, phonenumber string) userData {
	newUser := userData{
		firstName:   firstname,
		lastName:    lastname,
		dateofBirth: dob,
		married:     married,
		phoneNumber: phonenumber,
		createdDate: time.Now(),
	}
	fmt.Println(newUser)
	return newUser
}
