package helper

import (
	"fmt"
	"os/exec"
	"strings"
)

/*
kill process is running to continue with new process
*/
func KillProcess() {
	cmd := exec.Command("tasklist")
	output, _ := cmd.Output()
	lines := strings.Split(string(output), "\n")

	for _, line := range lines {
		if strings.Contains(line, "openvpn") {
			fields := strings.Fields(line)
			fmt.Println("fields", fields)
			pid := fields[1]
			fmt.Println("Killing process", pid)
			cmd := exec.Command("taskkill", "/F", "/PID", pid)
			cmd.Run()
		}
	}
}
