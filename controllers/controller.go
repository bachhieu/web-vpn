package controllers

import "os"

type VpnController struct {
}

type TestController struct {
}

var env = struct {
	VPNGATE string
}{
	os.Getenv("VPNGATE"),
}
