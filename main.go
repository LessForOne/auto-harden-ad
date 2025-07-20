package main

import (
	"auto-harden-ad/pkg/smb"
	"fmt"
)

func main() {
	fmt.Println("ForOne AD TOOL")
	fmt.Println("[1] - Desactiver SMBv1")
	fmt.Println("à venir")
	fmt.Println("à venir")
	fmt.Println("à venir")

	var choice int
	fmt.Println("Entrer un choix")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		smb.Smb()
		fmt.Println("smbv1 bien desactivé")
	case 2:
		fmt.Println("pas dispo")
	default:
		fmt.Println("non reconnu")
	}
}
