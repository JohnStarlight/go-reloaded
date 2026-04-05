package main

// BuildOutput reconstructs the final string from the transformed token slice.
// The main challenges here are spacing and quote handling — everything else is straightforward.
func BuildOutput(tokens []Token) string {
	result := ""
	inQuote := false      // are we currently inside a ' or " pair?
	firstInQuote := false // is the next word the first one inside a quote?
	first := true         // have we written anything to result yet?
	afterNewline := false // skip the leading space right after a newline
	for _, t := range tokens {
		switch t.Type {
		case Command:
			// commands were already applied in Transform, nothing to emit here

		case Punctuation:
			switch t.Value {
			case "'", "\"":
				// opening quote: add a space before it (unless it's the very first thing),
				// then set the flag so the first word inside gets no leading space.
				// closing quote: no space before, just close it out.
				switch inQuote {
				case false:
					switch first {
					case true:
						result += t.Value
						first = false
					default:
						result += " " + t.Value
					}
					inQuote = true
					firstInQuote = true
				case true:
					result += t.Value
					inQuote = false
				}
			default:
				// all other punctuation attaches directly to whatever came before it
				result += t.Value
			}

		case Newline:
			result += t.Value
			afterNewline = true

		case Word:
			switch {
			case inQuote:
				// first word after an opening quote has no space; subsequent ones do
				switch firstInQuote {
				case true:
					result += t.Value
					firstInQuote = false
				default:
					result += " " + t.Value
				}
			case first || afterNewline:
				// no leading space at the very start of the output or right after a newline
				result += t.Value
				first = false
				afterNewline = false
			default:
				result += " " + t.Value
			}
		default:
		}
	}
	return result
}
