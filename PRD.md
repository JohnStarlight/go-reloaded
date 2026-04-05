# go-reloaded PRD

## Overview
A text transformation tool written in Go that reads an input file, applies a series of modifications, and writes the result to an output file.

## Usage
```bash
go run . input.txt output.txt
```

## Features

### Number Base Conversion
- `(hex)` converts the previous word from hexadecimal to decimal
- `(bin)` converts the previous word from binary to decimal
- If the previous word is not a valid number, it is left unchanged

### Case Transformation
- `(up)` converts the previous word to uppercase
- `(low)` converts the previous word to lowercase
- `(cap)` capitalizes the first letter of the previous word
- `(up, N)`, `(low, N)`, `(cap, N)` apply the transformation to the N previous words
- Commands are case-insensitive and applied sequentially

### Punctuation Formatting
- Punctuation marks (`.`, `,`, `!`, `?`, `:`, `;`) are placed directly after the previous word with a space after them
- Groups of punctuation (`...`, `!?`) are kept together
- Single quotes `'` and double quotes `"` are placed directly around the quoted text without spaces

### Article Correction
- `a` is converted to `an` when the next word begins with a vowel (`a`, `e`, `i`, `o`, `u`) or `h`
- Case is preserved (`a` → `an`, `A` → `An`)
- Applied before case transformations to ensure correct results

## Technical Design

### Pipeline
```
Input file → Tokenizer → Vowels → Transform → BuildOutput → Output file
```

### Tokenizer
The input text is split into tokens of four types:
- **Word**: letters, digits, and hyphens
- **Punctuation**: punctuation marks grouped together
- **Command**: bracketed instructions like `(up)`, `(hex, 2)`
- **Newline**: line breaks, preserved as-is

### Transform
Iterates over tokens sequentially. When a Command token is found, it looks back through the token slice to find the N previous Word tokens and applies the appropriate transformation.

### Output Builder
Reconstructs the text from tokens:
- Words are separated by spaces
- Punctuation attaches directly to the previous word
- Newlines are preserved without trailing spaces
- Quoted text is wrapped without inner spaces

## Edge Cases Handled
- Command with no previous word → ignored
- Invalid hex/bin input → left unchanged
- Multiple newlines → preserved as-is
- Empty file → empty output
- Sequential commands on same word → applied in order
- Non-command parentheses like `(upper management)` → treated as regular word
- Hyphenated words like `well-known` → kept as single token
- First letter after non-letter character (e.g. `(not!)`) → correctly capitalized
