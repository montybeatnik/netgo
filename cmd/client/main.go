package main

import (
	"fmt"
	"log"
	"os"

	"github.com/montybeatnik/devcon"
)

var (
	homeLabIP = "10.0.0.60"
	mac       = "10.0.0.80"
	prodCisco = "172.30.83.76"
	localhost = "127.0.0.1"
)

func main() {
	un := os.Getenv("USER")
	pw := "password"
	// pw := os.Getenv("PASSWORD")
	client := devcon.NewClient(un,
		localhost,
		devcon.SetPassword(pw),
		devcon.SetPort("2222"),
	)
	output, err := client.Run("show version")
	if err != nil {
		fmt.Fprintf(os.Stderr, "command failed: %v\n", err)
		os.Exit(42)
	}
	fmt.Println(output)
	cmds := []string{
		"show ip interface brief",
		"exit",
	}
	interfaceOutput, err := client.RunAll(cmds...)
	if err != nil {
		log.Println("command failed", err)
		os.Exit(42)
	}
	fmt.Println(interfaceOutput)
}
