package helper

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
	"time"
)

/*
* @return bool
true if vpn live
false if vpn not live
*/
func CheckVpnIsLive() bool {
	configFile := "./config.ovpn"

	// run the OpenVPN client with the configuration file
	// cmd := exec.Command("openvpn", "--config", configFile)
	cmd := exec.Command("openvpn", "--config", configFile, "--data-ciphers", "AES-256-GCM:AES-128-GCM:AES-128-CBC")
	fmt.Printf("import file config and run cmd \n")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Start()           // run cmd
	time.Sleep(10 * time.Second) // await cmd run in 10s
	err = cmd.Process.Kill()     // close cmd
	if err != nil {
		fmt.Printf("\n Failed to kill process %s \n", err)

	}
	// parse the output to determine if the connection is successful
	output := stdout.String() + stderr.String()
	re := regexp.MustCompile(`Initialization Sequence Completed`)
	if re.MatchString(output) {
		_, err := exec.Command("ping", "8.8.8.8", "-w", "2").Output() // ping to 8.8.8.8 check vpn is live wait 2s
		if err != nil {
			return false
		} else {
			return true
		}
	} else {
		return false
	}
}
