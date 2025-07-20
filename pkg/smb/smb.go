package smb

import (
	"fmt"
	"os"
	"os/exec"
)

func Smb() error {
	fmt.Println("Test")

	cmd := exec.Command("powershell", "Set-SmbServerConfiguration", "-EnableSMB1Protocol", "$false")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("Erreur: ", err)
	}
	return nil
}
