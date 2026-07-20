package server

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"ubcNimAssignment/blueprint"
)

type Config struct {
	NimServerAddress string `json:"NimServerAddress"`
}

func main() {
	stateMoveMessage := blueprint.StateMoveMessage{
		GameState: []int8{0, 0, 0},
		MoveRow:   0,
		MoveCount: 0,
	}
	_ = stateMoveMessage

	EncodedconfigFile, err := os.ReadFile("config.json")
	if err != nil {
		fmt.Printf("Error reading config file: %v\n", err)
		return
	}

	var DecodedConfig Config
	if err := json.Unmarshal(EncodedconfigFile, &DecodedConfig); err != nil {
		fmt.Printf("Error decoding config file: %v\n", err)
		return
	}

	addr, err := net.ResolveUDPAddr("UDP", DecodedConfig.NimServerAddress)
	if err != nil {
		fmt.Printf("config file was not resolved completelty: %v\n", err)
		return
	}

	conn, err := net.ListenUDP("UDP", addr)
	if err != nil {
		fmt.Printf("config file was not resolved completelty: %v\n", err)
		return
	}

	defer conn.Close()

	fmt.Printf("Mock Nim Server listening on %s...\n", DecodedConfig.NimServerAddress)

	buffer := make([]byte, 1024)
	var currentBoard []uint8

	for{
		
	}
}
