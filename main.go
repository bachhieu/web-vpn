package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
)

func main() {
	// Đọc nội dung của tệp cấu hình OpenVPN
	configFile := "./vpngate_219.100.37.52_tcp_443.ovpn"
	// configFile := "./1.ovpn"

	// run the OpenVPN client with the configuration file
	cmd := exec.Command("openvpn", "--config", configFile)
	// cmd := exec.Command("echo", "hello")
	fmt.Printf("import file config 1 \n")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	fmt.Printf("import file config 2 \n")
	err := cmd.Start()
	fmt.Println("err :", err)
	if err == nil {
		fmt.Println("Failed to start OpenVPN:", err)
		fmt.Println("---------------------")

	} else {
		fmt.Println("Failed to start OpenVPN:", err)
		fmt.Println("++++++++++++++++++")
		return
	}

	// parse the output to determine if the connection is successful
	output := stdout.String() + stderr.String()
	fmt.Println("output: ", output)
	re := regexp.MustCompile(`Initialization Sequence Completed`)
	fmt.Println("\n re: ", re)
	fmt.Println("\n re.MatchString(output) ", re.MatchString(output))
	if re.MatchString(output) {
		fmt.Println("VPN is live")
		// do something with the live VPN connection here
	} else {
		fmt.Println("VPN is not live:", output)
		// handle the case where the VPN connection failed
	}

	// terminate the OpenVPN client
	err = cmd.Wait()
	if err != nil {
		fmt.Println("Failed to terminate OpenVPN:", err)
		return
	}
}
