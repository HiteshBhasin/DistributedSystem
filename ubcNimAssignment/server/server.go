package server

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"os"
	"ubcNimAssignment/blueprint"
)

type Config struct {
	NimServerAddress string `json:"NimServerAddress"`
}

func main() {
	stateMoveMessage := blueprint.StateMoveMessage{}

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

	for {

		n, remoteAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Printf("Error reading: %v\n", err)
			continue
		}

		err = json.Unmarshal(buffer[:n], &stateMoveMessage)
		if err != nil {
			fmt.Printf("Failed to unmarshal message: %v\n", err)
			continue
		}

		if stateMoveMessage.MoveCount == -1 {
			seed := int64(stateMoveMessage.MoveCount)
			fmt.Printf("Received game initialization with seed: %d\n", seed)

			r := rand.New(rand.NewSource(seed))
			row := r.Intn(3) + 3
			currentBoard = make([]uint8, row)

			for i := 0; i < row; i++ {
				currentBoard[i] = uint8(r.Intn(5) + 3)
			}

			gameStateInt8 := make([]int8, len(currentBoard))
			for i, v := range currentBoard {
				gameStateInt8[i] = int8(v)
			}

			reply := blueprint.StateMoveMessage{
				GameState: gameStateInt8,
				MoveCount: -1,
				MoveRow:   stateMoveMessage.MoveRow,
			}

			sendResponce(conn, remoteAddr, reply)
			printBoard(currentBoard)
			continue
		}

		fmt.Printf("Client moved: Removed %d from Row %d\n", stateMoveMessage.MoveCount, stateMoveMessage.MoveRow)

		currentBoard = int8BoardToUint8(stateMoveMessage.GameState)

		if isGameOver(currentBoard) {
			fmt.Println("Client made the last move. Client wins!")
			continue
		}

		// Server makes a simple, valid counter-move
		currentBoard = makeServerMove(currentBoard)

		if isGameOver(currentBoard) {
			fmt.Println("Server made the last move. Server wins!")
		}

		reply := blueprint.StateMoveMessage{
			GameState: uint8BoardToInt8(currentBoard),
			MoveRow:   0, // Mock details for tracking
			MoveCount: 1,
		}
		sendResponce(conn, remoteAddr, reply)
		printBoard(currentBoard)
	}

}

func makeServerMove(board []uint8) []uint8 {
	// Strategy: Find the first row with coins and take exactly 1 coin
	for i := 0; i < len(board); i++ {
		if board[i] > 0 {
			board[i]--
			fmt.Printf("Server responds: Removed 1 from Row %d\n", i)
			break
		}
	}
	return board
}

func isGameOver(board []uint8) bool {
	for _, coins := range board {
		if coins > 0 {
			return false
		}
	}
	return true
}
func sendResponce(conn *net.UDPConn, addr *net.UDPAddr, msg blueprint.StateMoveMessage) {
	data, _ := json.Marshal(msg)
	conn.WriteToUDP(data, addr)

}

func printBoard(board []uint8) {
	fmt.Printf("Current Board State: %v\n\n", board)
}

func int8BoardToUint8(board []int8) []uint8 {
	converted := make([]uint8, len(board))
	for i, v := range board {
		converted[i] = uint8(v)
	}
	return converted
}

func uint8BoardToInt8(board []uint8) []int8 {
	converted := make([]int8, len(board))
	for i, v := range board {
		converted[i] = int8(v)
	}
	return converted
}
