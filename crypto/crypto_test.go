package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash"
	"testing"
)

const (
	text       = "Hello test BEEVA"
	textMD5    = "45c780a8c42dfe4a55596e974b613cb1"
	textSHA1   = "523f71fb3878c8f686d9eeaee2f69ada185353ba"
	textSHA256 = "514625cc78ea0114fab2a696283c3ddc363f5ca3af98db88e0e9596cee711628"
	textSHA512 = "5479deecdb0901bca885bd8a84353ab7b7eaf4040ee550102066ac6c27c28f4140ef3aa3c857766f6ec8f6b803f89ea9a32743d528742e29d6bfc8474add9734"
	key        = "example key 1234"
)

func testSimpleChecker(expected, result string, t *testing.T) {
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

func testHashHelper(calculatedHash string, hash hash.Hash, t *testing.T) {
	result := Hash(text, hash)
	testSimpleChecker(calculatedHash, result, t)
}

func TestMD5(t *testing.T) {
	testHashHelper(textMD5, md5.New(), t)
}

func TestSHA1(t *testing.T) {
	testHashHelper(textSHA1, sha1.New(), t)
}

func TestSHA256(t *testing.T) {
	testHashHelper(textSHA256, sha256.New(), t)
}

func TestSHA512(t *testing.T) {
	testHashHelper(textSHA512, sha512.New(), t)
}

func TestAES(t *testing.T) {
	plaintextEncrypted := encryptAes(text, key)
	fmt.Printf("Encrypted text for '%s' string: '%s'\n", text, plaintextEncrypted)
	plaintext := decryptAes(plaintextEncrypted, key)
	fmt.Printf("Decrypted text for '%s' string: '%s'\n", plaintextEncrypted, plaintext)
	testSimpleChecker(text, plaintext, t)
}
