package main

import (
	"os"
	"bytes"
	"encoding/gob"
	"errors"
)

func playersToBinary(players []PlayerStore) ([]byte, error) {
	var buffer bytes.Buffer
	
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(players)

	if err != nil {
		return nil, errors.New("failed to encode players")
	}

	return buffer.Bytes(), nil
}

func binaryToPlayers(data []byte) ([]PlayerStore, error) {
	var players []PlayerStore

	buffer := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buffer)
	err := decoder.Decode(&players)

	if err != nil {
		return nil, errors.New("failed to decode player")
	}

	return players, nil
}

func openFile(path string) ([]byte, int64) {
	file, err := os.OpenFile(path, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fileInfo, _ := file.Stat()
	fileSize := fileInfo.Size()
	buffer := make([]byte, fileSize)

	file.Read(buffer)

	return buffer, fileSize
}

func writeFile(path string, data []byte) {
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.Write(data)
}

func GetPlayers() []PlayerStore {
	data, size := openFile("players.bin")

	if size == 0 {
		return []PlayerStore{}
	}

	players, err := binaryToPlayers(data)

	if err != nil {
		panic(err)
	}

	return players
}

func WritePlayers(playerStores []PlayerStore) {
	data, err := playersToBinary(playerStores)

	if err != nil {
		panic(err)
	}

	writeFile("players.bin", data)
}