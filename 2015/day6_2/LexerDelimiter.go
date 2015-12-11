package main

import "github.com/adampresley/lexer"

/*
LexerDelimiter processes delimiters. In this dialect delimiters are commas and the word
'through'.
*/
func LexerDelimiter(lex *lexer.Lexer) lexer.LexFn {
	_ = "breakpoint"

	/*
	 * We've reached the end of a line after a number. Move on to the next command.
	 */
	if lex.IsNewline() {
		return LexerCommand
	}

	lex.SkipWhitespace()
	delim := ""

	/*
	 * Continue until we get a comma or the word 'through'
	 */
	for {
		if lex.IsEOF() {
			lex.Emit(lexer.TOKEN_EOF)
			return lex.Errorf("Unexpected EOF. Expected a delimiter")
		}

		char := lex.CurrentCharacter()

		if lex.IsWhitespace() {
			lex.Emit(TOKEN_DELIMITER)
			return LexerCommand
		}

		if char == "," {
			lex.Inc(1)
			lex.Emit(TOKEN_DELIMITER)
			return LexerNumber
		}

		delim += lex.CurrentCharacter()
		if delim == "through" {
			lex.Inc(1)
			lex.Emit(TOKEN_DELIMITER)
			return LexerNumber
		}

		lex.Inc(1)
	}
}
