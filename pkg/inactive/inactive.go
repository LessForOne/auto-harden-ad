package inactive

import (
	"fmt"
	"os"
	"os/exec"
)

func Inact() error {
	fmt.Println("Test")

	cmd := exec.Command("powershell", "Search-ADAccount -AccountInactive -UsersOnly | Disable-ADAccount")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("Erreur: ", err)
	}
	return nil
}
