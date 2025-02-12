package main

import (
	"os"
	"reflect"
	"testing"
)

func TestPlayersBinaryConversion(t *testing.T) {
	players := []PlayerStore{
		{"Player1", 100},
		{"Player2", 200},
	}

	// Test playersToBinary
	data, err := playersToBinary(players)
	if err != nil {
		t.Errorf("playersToBinary failed: %v", err)
	}

	// Test binaryToPlayers
	decodedPlayers, err := binaryToPlayers(data)
	if err != nil {
		t.Errorf("binaryToPlayers failed: %v", err)
	}

	if !reflect.DeepEqual(players, decodedPlayers) {
		t.Errorf("Players not equal after encoding/decoding.\nExpected: %v\nGot: %v", players, decodedPlayers)
	}
}

func TestFileOperations(t *testing.T) {
	testFile := "test.bin"
	testData := []byte("test data")

	// Test writeFile
	writeFile(testFile, testData)

	// Test openFile
	data, size := openFile(testFile)
	if size != int64(len(testData)) {
		t.Errorf("File size mismatch. Expected: %d, Got: %d", len(testData), size)
	}

	if !reflect.DeepEqual(data, testData) {
		t.Errorf("File data mismatch.\nExpected: %v\nGot: %v", testData, data)
	}

	// Cleanup
	os.Remove(testFile)
}

func TestPlayersStorage(t *testing.T) {
	players := []PlayerStore{
		{"Player1", 100},
		{"Player2", 200},
	}

	WritePlayers(players)
	loadedPlayers := GetPlayers()

	if !reflect.DeepEqual(players, loadedPlayers) {
		t.Errorf("Players not equal after save/load.\nExpected: %v\nGot: %v", players, loadedPlayers)
	}

	// Cleanup
	os.Remove("players.bin")
}
