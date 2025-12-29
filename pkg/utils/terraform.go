package utils

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

func EndsWithTf(str string) bool {
	return strings.HasSuffix(str, ".tf")
}

func RandomName() string {
	randomBytes := make([]byte, 5)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}
	randomString := base64.RawURLEncoding.EncodeToString(randomBytes)
	return fmt.Sprintf("terraform-%s.tf", randomString)
}

func GetName(name string) string {
	name = RemoveBlankLinesFromString(name)
	if EndsWithTf(name) {
		return name
	}
	return RandomName()
}

func TerraformPath() (string, error) {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("where", "terraform")
	} else {
		cmd = exec.Command("which", "terraform")
	}
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("error running Init:%w", err)
	}
	return strings.TrimRight(string(output), "\n"), nil
}
