# go-reloaded

A text transformation tool that reads an input file, applies a set of formatting rules and inline commands, and writes the result to an output file.

## Usage

```sh
go run . <input_file> <output_file>
```

**Example:**

```sh
go run . test.txt result.txt
```

## Transformations

The tool processes text tokens and applies the following rules:

### Inline Commands

Place a command immediately after the word(s) it should affect:

| Command | Effect |
|---|---|
| `(hex)` | Converts the preceding word from hexadecimal to decimal |
| `(bin)` | Converts the preceding word from binary to decimal |
| `(up)` | Uppercases the preceding word |
| `(low)` | Lowercases the preceding word |
| `(cap)` | Capitalizes the preceding word |
| `(up, N)` | Uppercases the preceding N words |
| `(low, N)` | Lowercases the preceding N words |
| `(cap, N)` | Capitalizes the preceding N words |

### Punctuation Spacing

Punctuation is attached to the word before it (no space before, space after):

```
Input:  kinda boring ,what do you think ?
Output: kinda boring, what do you think?
```

Ellipses and multi-character punctuation are treated as a single unit:

```
Input:  are ... kinda
Output: are... kinda
```

### Quote Handling

Single (`'`) and double (`"`) quotes wrap their content without extra spaces:

```
Input:  She said " hello world "
Output: She said "hello world"
```

### Article "a" / "an" Correction

The article `a` is automatically changed to `an` when the next word starts with a vowel or `h` (for the purposes of this task, we do not check wether the `h` is silent or not):

```
Input:  I saw a elephant and a honest man
Output: I saw an elephant and an honest man
```

## Input/Output Example

**Input (`test.txt`):**
```
Punctuation tests are ... kinda boring ,what do you think ?
```

**Output:**
```
Punctuation tests are... kinda boring, what do you think?
```

## Running Tests

The project has no external dependencies. Run directly with the Go toolchain (Go 1.21+):

```sh
go run . input.txt output.txt
```
