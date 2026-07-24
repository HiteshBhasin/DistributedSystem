package client

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"ubcNimAssignment/blueprint"
)

type Config struct {
	ServerAddress string `json:"server UDP ADD"`
}

var config = Config{
	ServerAddress: "197.0.0.1:8080",
}

func main() {
	byteData := ReadConfig("config.json")
	msg, err := WriteByteFile(byteData)
	if err != nil {
		fmt.Printf("Error parsing config: %v\n", err)
		return
	}
	fmt.Printf("Loaded message: %+v\n", msg)

	serverAddr, err := net.ResolveUDPAddr("udp", config.ServerAddress)
	if err != nil {
		fmt.Printf("Error resolving UDP address: %v\n", err)
		return
	}
	conn, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		fmt.Printf("Error dialing UDP: %v\n", err)
		return
	}
	defer conn.Close()
}

func ReadConfig(configFile string) []byte {
	fileBytes, err := os.ReadFile(configFile)
	if err != nil {
		fmt.Printf("Error reading config file: %v\n", err)

	}
	return fileBytes
}

func WriteByteFile(data []byte) (blueprint.StateMoveMessage, error) {
	var msg blueprint.StateMoveMessage
	if err := json.Unmarshal(data, &msg); err != nil {
		return msg, err
	}
	return msg, nil
}
