package main

import (
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {
	// Run the command inside the auth-api folder
	authCmd := exec.Command("go", "run", ".")
	authCmd.Dir = "./auth-api"
	authCmd.Stdout = os.Stdout
	authCmd.Stderr = os.Stderr

	// Run the command inside the user-api folder concurrently
	userCmd := exec.Command("go", "run", ".")
	userCmd.Dir = "./user-api"
	userCmd.Stdout = os.Stdout
	userCmd.Stderr = os.Stderr

	// Start executing both commands concurrently
	go func() {
		if err := authCmd.Run(); err != nil {
			log.Fatal(err)
			authCmd.Process.Kill()
		}
	}()

	go func() {
		if err := userCmd.Run(); err != nil {
			log.Fatal(err)
			userCmd.Process.Kill()
		}
	}()

	for {
		time.Sleep(time.Second) // Introduce a delay to prevent 100% CPU usage
	}
}
