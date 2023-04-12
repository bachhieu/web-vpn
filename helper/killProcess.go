package helper

import (
	"fmt"
	"os/exec"
	"strings"
)

func KillProcess() {
	cmd := exec.Command("tasklist")
	output, _ := cmd.Output()
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.Contains(line, "openvpn") {
			fields := strings.Fields(line)
			pid := fields[1]
			fmt.Println("Killing process", pid)
			cmd := exec.Command("taskkill", "/F", "/PID", pid)
			cmd.Run()
		}
	}
}