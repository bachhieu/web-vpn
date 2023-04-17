package utils

import (
	"log"
	"os/exec"
)

func Init() {
	// Create directory /dev/net
	cmd := exec.Command("mkdir", "-p", "/dev/net")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	// Create TUN deexec
	cmd = exec.Command("mknod", "/dev/net/tun", "c", "10", "200")
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	// Set permissions on TUN device
	cmd = exec.Command("chmod", "600", "/dev/net/tun")
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	// mkdir -p /dev/net && \
	// mknod /dev/net/tun c 10 200 && \
	// chmod 600 /dev/net/tun
	log.Println("TUN device created successfully.")
}
