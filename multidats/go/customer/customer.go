package customer

import (
	"fmt"
)

func (user CustomerData) Generatedata() *CustomerData {
	fmt.Println(&user)
	return &user
}

func (user CustomerData) FindTag(tag string) bool {
	status := false
	for key, value := range user.Tags {
		if value == tag {
			fmt.Println("key:", key)
			fmt.Println("value:", value)
			status = true
			break
		}
	}
	fmt.Println(status)
	return status
}
