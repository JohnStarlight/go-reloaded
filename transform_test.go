package main

import "testing"

func TestHex(t *testing.T) {
	tokens := Tokenizer("1E (hex) files")
	tokens = Transform(tokens)
	result := BuildOutput(tokens)
	if result != "30 files" {
		t.Errorf("Expected '30 files', got '%s'", result)
	}
}

func TestBin(t *testing.T) {
	tokens := Tokenizer("10 (bin) years")
	tokens = Transform(tokens)
	result := BuildOutput(tokens)
	if result != "2 years" {
		t.Errorf("Expected '2 years', got '%s'", result)
	}
}

func TestUp(t *testing.T) {
	tokens := Tokenizer("hello world (up, 2)")
	tokens = Transform(tokens)
	result := BuildOutput(tokens)
	if result != "HELLO WORLD" {
		t.Errorf("Expected 'HELLO WORLD', got '%s'", result)
	}
}

func TestLow(t *testing.T) {
	tokens := Tokenizer("HELLO WORLD (low, 2)")
	tokens = Transform(tokens)
	result := BuildOutput(tokens)
	if result != "hello world" {
		t.Errorf("Expected 'hello world', got '%s'", result)
	}
}

func TestCap(t *testing.T) {
	tokens := Tokenizer("hello world (cap, 2)")
	tokens = Transform(tokens)
	result := BuildOutput(tokens)
	if result != "Hello World" {
		t.Errorf("Expected 'Hello World', got '%s'", result)
	}
}

func TestVowels(t *testing.T) {
	tokens := Tokenizer("a elephant walks")
	tokens = Vowels(tokens)
	tokens = Transform(tokens)
	result := BuildOutput(tokens)
	if result != "an elephant walks" {
		t.Errorf("Expected 'an elephant walks', got '%s'", result)
	}
}

func TestPunctuation(t *testing.T) {
	tokens := Tokenizer("hello , world !")
	tokens = Transform(tokens)
	result := BuildOutput(tokens)
	if result != "hello, world!" {
		t.Errorf("Expected 'hello, world!', got '%s'", result)
	}
}

func TestQuotes(t *testing.T) {
	tokens := Tokenizer("' hello world ' how are you")
	tokens = Transform(tokens)
	result := BuildOutput(tokens)
	if result != "'hello world' how are you" {
		t.Errorf("Expected \"'hello world' how are you\", got '%s'", result)
	}
}
