package main

import (
	"strconv"
	"strings"
	"unicode"
)

// capitalize uppercases just the first letter of a word, leaving the rest alone.
// scans for the first actual letter so it works on things like "(not!)" too.
func capitalize(word string) string {
	runes := []rune(word)
	for i, r := range runes {
		if unicode.IsLetter(r) {
			runes[i] = unicode.ToUpper(r)
			return string(runes)
		}
	}
	return word
}

// Transform applies all the Command tokens to the words around them.
// Vowels is called first so that a/an correction happens before any case changes.
func Transform(tokens []Token) []Token {
	tokens = Vowels(tokens)
	for i, t := range tokens {
		if t.Type == Command {
			switch {
			case t.Value == "(hex)":
				// parse the previous word as base-16 and replace it with the decimal value
				j := findPrev(tokens, i)
				if j == -1 {
					break
				}
				tmp := tokens[j].Value
				res, err := strconv.ParseInt(tmp, 16, 64)
				if err != nil {
					break
				}
				tokens[j].Value = strconv.Itoa(int(res))
			case t.Value == "(bin)":
				// same idea but base-2
				j := findPrev(tokens, i)
				if j == -1 {
					break
				}
				tmp := tokens[j].Value
				res, err := strconv.ParseInt(tmp, 2, 64)
				if err != nil {
					break
				}
				tokens[j].Value = strconv.Itoa(int(res))
			case strings.HasPrefix(t.Value, "(up"):
				// walk back Count words and uppercase each one
				j := i
				for k := 0; k < t.Count; k++ {
					j = findPrev(tokens, j)
					if j == -1 {
						break
					}
					tokens[j].Value = strings.ToUpper(tokens[j].Value)
				}
			case strings.HasPrefix(t.Value, "(low"):
				// same but lowercase
				j := i
				for k := 0; k < t.Count; k++ {
					j = findPrev(tokens, j)
					if j == -1 {
						break
					}
					tokens[j].Value = strings.ToLower(tokens[j].Value)
				}
			case strings.HasPrefix(t.Value, "(cap"):
				// capitalize the first letter of each of the Count previous words
				j := i
				for k := 0; k < t.Count; k++ {
					j = findPrev(tokens, j)
					if j == -1 {
						break
					}
					tokens[j].Value = capitalize(tokens[j].Value)
				}
			}
		}
	}
	return tokens
}
