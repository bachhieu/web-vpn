package helper

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"syscall"
	"time"
)

/*
* @return bool
true if vpn live
false if vpn not live
*/
func CheckVpnIsLive(content []byte) bool {
	//  Create a temporary file to write the configuration
	tmpfile, err := ioutil.TempFile("./", "*.ovpn")
	if err != nil {
		return false
	}

	// clean file configuration after check vpn
	defer os.Remove(tmpfile.Name()) // clean up

	if _, err := tmpfile.Write(content); err != nil {
		return false
	}
	if err := tmpfile.Close(); err != nil {
		return false
	}

	// run the OpenVPN client with the configuration file
	// cmd := exec.Command("openvpn", "--config", tmpfile.Name())
	cmd := exec.Command("openvpn", "--config", tmpfile.Name(), "--data-ciphers", "AES-256-GCM:AES-128-GCM:AES-128-CBC")
	// cmd := exec.Command("openvpn", "--config", "./config.ovpn", "--data-ciphers", "AES-256-GCM:AES-128-GCM:AES-128-CBC")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err = cmd.Start()                         // run cmd
	time.Sleep(10 * time.Second)              // await cmd run in 10s
	err = cmd.Process.Signal(syscall.SIGTERM) // close cmd
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
