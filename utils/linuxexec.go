package utils

import (
	"fmt"
	"os/exec"
	"strings"
)

func Execute(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to execute '%s %s': %v", command, strings.Join(args, " "), err)
	}
	return strings.TrimSpace(string(output)), nil
}