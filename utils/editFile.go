package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
path is a path to file want edit,
content is new content or new line,
overwire is string want overwrite.
*/
func EditFile(path string, content []byte, overwire string) error {

	// Mở file .env để đọc
	envFile, err := os.Open(path)
	if err != nil {
		return err
	}
	defer envFile.Close()

	// Tạo một slice chứa các dòng trong file .env
	var envLines []string
	scanner := bufio.NewScanner(envFile)
	for scanner.Scan() {
		envLines = append(envLines, scanner.Text())
	}

	// Tìm và thay đổi giá trị của trường "AUTO"
	for i, line := range envLines {
		if strings.HasPrefix(line, overwire) {
			envLines[i] = string(content)
			break
		}
	}

	// Mở file .env để ghi và ghi lại các dòng đã được thay đổi
	envFile, err = os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer envFile.Close()

	writer := bufio.NewWriter(envFile)
	for _, line := range envLines {
		fmt.Fprintln(writer, line)
	}
	writer.Flush()
	return nil
}
