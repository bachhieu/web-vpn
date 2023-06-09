// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import (
	"io/ioutil"

	"github.com/swaggo/swag"
)

var data, _ = ioutil.ReadFile("./docs/swagger.yaml")
var docTemplate = string(data) ;

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
    Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Swagger Example API",
	Description:      "This is a sample server Petstore server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,   
    
}


func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
