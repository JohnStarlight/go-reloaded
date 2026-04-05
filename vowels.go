package main

import "unicode"

// firstLetter returns the first Unicode letter in a word, or 0 if there isn't one.
// needed because some tokens start with punctuation (e.g. quotes).
func firstLetter(word string) rune {
	for _, r := range word {
		if unicode.IsLetter(r) {
			return r
		}
	}
	return 0
}

// Vowels fixes "a" → "an" when the next word starts with a vowel or h.
// we don't check whether the h is actually silent, we just always use "an" before h.
func Vowels(tokens []Token) []Token {
	for i := 0; i < len(tokens); i++ {
		switch tokens[i].Type {
		case Word:
			switch tokens[i].Value {
			case "a", "A":
				j := findNext(tokens, i)
				if j == -1 {
					break
				}
				first := firstLetter(tokens[j].Value)
				switch first {
				case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U', 'h', 'H':
					tokens[i].Value += "n" // "a" → "an", "A" → "An"
				}
			}
		}
	}
	return tokens
}
