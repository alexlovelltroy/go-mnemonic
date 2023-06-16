package mnemonic

import (
	"strings"
	"testing"
)

func TestGetRandomName(t *testing.T) {

	// Set up test data
	separator := "-"

	// Test case 1: length = 1
	name := GetRandomName(1, separator)
	if !contains(wordlist, name) {
		t.Errorf("Expected %s, but got %s", wordlist[0], name)
	}

	// Test case 2: length = 2
	name = GetRandomName(2, separator)
	parts := strings.Split(name, separator)
	if len(parts) != 2 {
		t.Errorf("Expected name to contain two words")
	}
	if !contains(wordlist, parts[0]) || !contains(wordlist, parts[1]) {
		t.Errorf("Expected name to contain two words from the wordlist")
	}

	// Test case 3: length = 5
	name = GetRandomName(5, separator)
	parts = strings.Split(name, separator)
	if len(parts) != 5 {
		t.Errorf("Expected name to contain five words")
	}
	for _, part := range parts {
		if !contains(wordlist, part) {
			t.Errorf("Expected word from the wordlist, but got %s", part)
		}
	}
}

func TestGetRandomWords(t *testing.T) {

	// Test case 1: length = 1
	words := GetRandomWords(1)
	if len(words) != 1 || !contains(wordlist, words[0]) {
		t.Errorf("Expected a single word from the wordlist")
	}

	// Test case 2: length = 5
	words = GetRandomWords(5)
	if len(words) != 5 {
		t.Errorf("Expected the length of words to be 5")
	}
	for _, word := range words {
		if !contains(wordlist, word) {
			t.Errorf("Expected word from the wordlist, but got %s", word)
		}
	}
}

// Helper function to check if a string is present in a slice of strings
func contains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}
