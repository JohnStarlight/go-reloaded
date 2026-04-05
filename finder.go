package main

// findPrev returns the index of the nearest Word token before position i,
// skipping over anything that isn't a Word (punctuation, commands, newlines).
// returns -1 if there's no Word token before i.
func findPrev(tokens []Token, i int) int {
	j := i - 1
	for j >= 0 && tokens[j].Type != Word {
		j--
	}
	return j
}

// findNext returns the index of the nearest Word token after position i.
// returns -1 if there's nothing left.
func findNext(tokens []Token, i int) int {
	j := i + 1
	for j < len(tokens) && tokens[j].Type != Word {
		j++
	}
	if j >= len(tokens) {
		return -1
	}
	return j
}
