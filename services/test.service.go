package services

import (
	"fmt"
)

func (ctl *TestService) Index() string {
	fmt.Println("<b>Thank you! " + "example" + "</b>")
	return "<b>Thank you! " + "example" + "</b>"

}
