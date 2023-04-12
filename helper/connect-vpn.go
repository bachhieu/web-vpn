package helper

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
	"time"
)

func CheckVpnIsLive() bool {
	// Đọc nội dung của tệp cấu hình OpenVPN
	configFile := "./vpn.ovpn"
	// configFile := "./1.ovpn"

	// run the OpenVPN client with the configuration file
	cmd := exec.Command("openvpn", "--config", configFile)
	// cmd := exec.Command("echo", "hello")
	fmt.Printf("import file config 1 \n")
	var stdout, stderr bytes.Buffer
	var err error
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	go func(){
		err = cmd.Start()

	}()
	time.Sleep(10 * time.Second)
	if err != nil {
		fmt.Println("Failed to start OpenVPN:", err)
		_,err = fmt.Println("Failed to start OpenVPN:")
		panic(err)

	} 
	
	// parse the output to determine if the connection is successful
	output := stdout.String() + stderr.String()
	fmt.Println(output)
	re := regexp.MustCompile(`Initialization Sequence Completed`)
	fmt.Println("\n re.MatchString(output) ", re.MatchString(output))
	if re.MatchString(output) {
		fmt.Println("VPN is live")
		return true
		// do something with the live VPN connection here
	} else {
		fmt.Println("VPN is not live:", output)
		return false
		// handle the case where the VPN connection failed
	}
	// // terminate the OpenVPN client
	// err = cmd.Wait()
	// if err != nil {
	// 	fmt.Println("Không có lỗi:")
	// 	return true
	// } else {
	// 	fmt.Println("có lỗi:")
	// 	return true
	// }
}
