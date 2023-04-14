package services

import (
	"fmt"
)

func (ctl *VpnService) Index() string {
	fmt.Println("<b>Thank you! " + "example" + "</b>")
	return "<b>Thank you! " + "example" + "</b>"

}
