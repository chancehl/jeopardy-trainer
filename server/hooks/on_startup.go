package hooks

import (
	"log"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func OnBeforeServerStart(mode string) {
	if mode == gin.ReleaseMode {
		cmd := exec.Command("npm", "run", "build")

		cmd.Dir = "./web"
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()

		if err != nil {
			log.Fatalf("Failed to execute command: %s\n", err)
		}
	}
}
