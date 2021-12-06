package wvexempt

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/brad-jones/goexec/v2"
	"github.com/brad-jones/winsudo"
)

// GetExemtions lists all the loopback exemptions.
func GetExemptions() (string, error) {
	Args := []string{
		"CheckNetIsolation.exe",
		"LoopbackExempt",
		"-s",
	}
	out, err := exec.Command(Args[0], Args[1:]...).Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

// CheckExempt determines if WebView is exempt from the Loopback restriction yet.
func CheckExempt() bool {
	exemptions, err := GetExemptions()
	if err != nil {
		return false
	}
	if strings.Contains(exemptions, "cw5n1h2txyewy") {
		return true
	}
	return false
}

// Exempt ensures that WebView is added to the Loopback Exemption list
func Exempt() error {
	if CheckExempt() {
		return nil
	}
	Args := []string{
		"CheckNetIsolation.exe",
		"LoopbackExempt",
		"-a",
		"-n=\"Microsoft.Win32WebViewHost_cw5n1h2txyewy\"",
	}

	// Execute the given command in an elevated environment
	exit, err := winsudo.ElevatedExec(Args[0], goexec.Args(Args[1:]...))
	if exit != 0 {
		err = fmt.Errorf("Elevation Error: WebView Fixer exited status %s", exit)
	}

	return err
}
