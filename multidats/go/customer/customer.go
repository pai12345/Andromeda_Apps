package customer

import (
	"fmt"
)

func (user CustomerData) Generatedata() *CustomerData {
	fmt.Println(&user)
	return &user
}
