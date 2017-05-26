package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
)

// FileChunk - chunk from the file with fixed length
type FileChunk struct {
	Content []byte
}

// CreateFileHash - create hash sum for the given file
func CreateFileHash(filename string) ([32]byte, error) {
	// read file
	content, err := ReadFile(filename)
	if err != nil {
		panic(err.Error())
	}

	// split file content into small chunks (1KB)
	chunks := SplitContentIntoChunks(content, 1024)

	fmt.Println(len(chunks))
	// fmt.Println(chunks[len(chunks)-1])
	// start from last chunk
	var hash []byte
	// set hash value as N - 1 item
	hash = chunks[len(chunks)-1]
	for i := len(chunks) - 2; i >= 0; i-- {
		// get SHA256 sum for current block
		sum := sha256.Sum256(hash)
		// add it to current block
		hash = append(chunks[i], sum[:]...)
	}

	// get final hash
	checksum := sha256.Sum256(hash)

	// checksum := hash[:]
	//

	fmt.Println(hex.EncodeToString(checksum[:]))
	return checksum, nil
}

// SplitContentIntoChunks - splint source file into chunks with limited size
func SplitContentIntoChunks(source []byte, limit int) [][]byte {
	fmt.Println("Source file length (bytes): " + strconv.Itoa(len(source)))
	var chunk []byte
	chunks := make([][]byte, 0, len(source)/limit+1)
	fmt.Println("Count chunks: " + strconv.Itoa(cap(chunks)))
	for i := 0; i < cap(chunks); i++ {
		if len(source) < limit {
			break
		}
		chunk, source = source[:limit], source[limit:]
		chunks = append(chunks, chunk)
	}
	if len(source) > 0 {
		chunks = append(chunks, source[:])
	}
	return chunks
}
