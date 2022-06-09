package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/montybeatnik/devcon"
)

func main() {
	client := devcon.NewClient(
		os.Getenv("SSH_USER"),
		"10.0.0.60",
		devcon.Password(os.Getenv("SSH_PASSWORD")),
		devcon.Timeout(time.Second*1),
	)
	out, err := client.Run("show version")
	if err != nil {
		log.Fatalf("command failed: %v", err)
	}
	fmt.Println(out)
}
