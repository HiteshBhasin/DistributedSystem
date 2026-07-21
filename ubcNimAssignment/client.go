package client

import (
	"fmt"
	"os"
	"ubcNimAssignment/blueprint"
)

func main() {
	msg = blueprint.StateMoveMessage
}

func ReadConfig(configFile string) []byte {
	fileBytes, err := os.ReadFile(configFile)
	if err != nil {
		fmt.Printf("Error reading config file: %v\n", err)

	}
	return fileBytes
}
