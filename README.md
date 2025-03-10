# Go Reloaded

Advanced Go programming exercises and implementations.

## Description

This project is a text processing tool that performs various transformations on input text files. It reads from an input file, applies a series of text editing operations, and writes the result to an output file.

## Files and Functions

1. **`main.go`**
    - **`main()`**: The main function. It checks command line arguments, reads the content of the input file, applies text editing operations, and writes the edited text to the output file.

2. **`HexBin.go`**
    - **`EditHexBin(text string) string`**: This function replaces hexadecimal and binary representations in the input text with their decimal equivalents.

3. **`Case.go`**
    - **`EditCase(text string) string`**: This function performs uppercase, lowercase, and capitalized transformations on the input text based on specified patterns.
    - **`processMultipleCase(text string, Trim *regexp.Regexp, caseFunc func(string) string) string`**: This function handles case modifications for patterns with a specific number of words to modify.

4. **`Punctuation.go`**
    - **`EditPunctuation(text string) string`**: This function handles punctuation marks in the input text, ensuring proper spacing and placement.

5. **`Quote.go`**
    - **`EditQuote(text string) string`**: This function cleans up quotation marks and ensures proper spacing around them.

6. **`AN.go`**
    - **`EditAN(text string) string`**: This function replaces occurrences of "a" with "an" before words starting with a vowel or 'h'.

## Usage

```bash
go run . input_file.txt output_file.txt
```

Example:
```
Input: This is (hex)FF in hexadecimal and (bin)1010 in binary.
Output: This is 255 in hexadecimal and 10 in binary.
```

## Features

- Conversion of hexadecimal and binary to decimal
- Text case transformations (uppercase, lowercase, capitalization)
- Proper punctuation handling and spacing
- Proper article usage ("a" vs "an")
- Quote formatting

## Implementation Details

- Regular expression-based text processing
- Modular code organization with separate files for different transformations
- Clean and efficient string manipulation
- Command-line argument handling

## Requirements

- Go 1.13 or higher
