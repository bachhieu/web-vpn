package services

import (
	"bachhieu/web-vpn/models"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func (ctl *VpnService) FindVpnlive() models.VpnModel {
	var vpn = models.GetCollection("vpn")
	  fmt.Printf("%v\n",vpn)

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	cur, err := vpn.Find(ctx, bson.M{"live": true})
	defer cur.Close(context.Background())
	vpnModel := models.VpnModel{}
	err = cur.Decode(&vpnModel)
	  if err != nil { log.Fatal(err) }
	  // do something with result...
	
	  // To get the raw bson bytes use cursor.Current
	//   fmt.Printf("%v\n",person.Id.String())
	fmt.Println("<b>Thank you! " + "example" + "</b>")
	return vpnModel

}
