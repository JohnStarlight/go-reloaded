package main

import (
	"strconv"
	"strings"
	"unicode"
)

type TokenType int

const (
	Word        TokenType = iota // 0
	Punctuation TokenType = iota // 1
	Command     TokenType = iota // 2
	Newline     TokenType = iota // 3
)

// Token is the basic unit we work with throughout the pipeline.
// Count is only meaningful for commands like (up, 3) — otherwise it's just 1.
type Token struct {
	Value string    // the collection of chars
	Type  TokenType // the label we put on each collection of chars (as the types specified above)
	Count int       // this is specific for the bracketed commands in the given text
}

// extractCount pulls the number out of a command like "(up, 3)" → 3.
// walks digit by digit so it handles multi-digit counts correctly.
func extractCount(command string) int {
	res := 0
	for _, v := range command {
		if unicode.IsDigit(v) {
			tmp, _ := strconv.Atoi(string(v))
			res = res*10 + tmp
		}
	}
	return res
}

// Tokenizer splits raw text into a flat slice of tokens.
// spaces are discarded (they get rebuilt by BuildOutput), everything else
// becomes a Word, Punctuation, Command, or Newline token.
func Tokenizer(text string) []Token {
	runes := []rune(text)
	tokens := []Token{}
	for i := 0; i < len(runes); i++ {
		switch {
		case runes[i] == '(':
			// scan ahead to the closing paren
			j := i
			for j < len(runes) && runes[j] != ')' {
				j++
			}
			command := string(runes[i : j+1])
			i = j
			cmdLower := strings.ToLower(command)
			// only treat it as a Command if it matches one of our known instructions,
			// otherwise fall through and keep it as a regular Word token
			if cmdLower == "(hex)" ||
				cmdLower == "(bin)" ||
				cmdLower == "(up)" ||
				cmdLower == "(low)" ||
				cmdLower == "(cap)" ||
				strings.HasPrefix(cmdLower, "(up, ") ||
				strings.HasPrefix(cmdLower, "(low, ") ||
				strings.HasPrefix(cmdLower, "(cap, ") {
				count := extractCount(command)
				if count == 0 {
					count = 1 // bare commands like (up) affect exactly one word
				}
				tokens = append(tokens, Token{Value: cmdLower, Type: Command, Count: count})
			} else {
				tokens = append(tokens, Token{Value: command, Type: Word})
			}
		case runes[i] == '\n': // newline
			tokens = append(tokens, Token{Value: "\n", Type: Newline})
		case runes[i] == ' ': // spaces are just separators, drop them
		case unicode.IsLetter(runes[i]) || unicode.IsDigit(runes[i]): // word
			// consume the whole word in one shot, hyphens included (e.g. "well-known")
			j := i
			for j < len(runes) && (unicode.IsLetter(runes[j]) || unicode.IsDigit(runes[j]) || runes[j] == '-') {
				j++
			}
			word := string(runes[i:j])
			tokens = append(tokens, Token{Value: word, Type: Word}) // add token
			i = j - 1
		default: // punctuation
			// greedily grab consecutive punctuation chars (handles "...", "!?", etc.)
			j := i
			for j < len(runes) &&
				!unicode.IsLetter(runes[j]) &&
				!unicode.IsDigit(runes[j]) &&
				runes[j] != ' ' &&
				runes[j] != '(' &&
				runes[j] != '\n' {
				j++
			}
			punctuation := string(runes[i:j])
			tokens = append(tokens, Token{Value: punctuation, Type: Punctuation})
			i = j - 1
		}
	}
	return tokens
}
