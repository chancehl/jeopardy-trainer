package hooks

import (
	"log"
	"os"
	"os/exec"
	"strconv"
)

func CheckIfShouldSkipBuild() bool {
	skipNpmBuildEnvVariable := os.Getenv("SKIP_NPM_BUILD")

	if skipNpmBuildEnvVariable == "" {
		return false
	}

	shouldBuild, boolConversionErr := strconv.ParseBool(skipNpmBuildEnvVariable)

	if boolConversionErr != nil {
		log.Fatalf("Error parsing the boolean value: %s\n", boolConversionErr)
	}

	return shouldBuild
}

func OnServerStart() {
	shouldSkipBuild := CheckIfShouldSkipBuild()

	if !shouldSkipBuild {
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
