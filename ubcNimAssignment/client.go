package client

import (
	"encoding/json"
	"fmt"
	"os"
	"ubcNimAssignment/blueprint"
)

type Config struct {
	ServerAddress string `json:"server UDP ADD"`
}

var config = Config{
	ServerAddress: "197.0.0.8080",
}

func main() {
	msg := blueprint.StateMoveMessage{}
	byteData = ReadConfig("config.json")

}

func ReadConfig(configFile string) []byte {
	fileBytes, err := os.ReadFile(configFile)
	if err != nil {
		fmt.Printf("Error reading config file: %v\n", err)

	}
	return fileBytes
}

func WriteByteFile(data []byte, msg blueprint.StateMoveMessage) {
	err := json.Unmarshal(data, &msg)
	if err != nil {
		fmt.Printf("successfully generated the file and returned the %v\n", err)
	}
}
