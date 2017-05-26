package main

import (
	"encoding/hex"
	"testing"
)

func TestHashMP4(t *testing.T) {
	h0 := "03c08f4ee0b576fe319338139c045c89c3e8e9409633bea29442e21425006ea8"
	relativeFilePath := "./samples/test.mp4"

	h0Test, err := CreateFileHash(relativeFilePath)
	if err != nil {
		t.Fatalf("Error while reading file %q", relativeFilePath)
	}

	hHex := hex.EncodeToString(h0Test[:])
	if h0 != hHex {
		t.Fatalf("Checksums doesn't match: expected=%q, real=%q", h0, hHex)
	}
}
