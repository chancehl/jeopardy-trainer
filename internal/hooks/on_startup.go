package hooks

import (
	"log"
	"os"
	"os/exec"
)

func OnServerStart() {
	cmd := exec.Command("npm", "run", "build")

	cmd.Dir = "./web"
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {
		log.Fatalf("Failed to execute command: %s\n", err)
	}
}
