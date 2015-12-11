package main

import "github.com/adampresley/lexer"

func LexerEscape(lex *lexer.Lexer) lexer.LexFn {
	_ = "breakpoint" // LexerEscape

	/*
	 * We start at a slash. Increment one. If that character is an 'x' then
	 * we have a Hex sequence. Otherwise the single character will do.
	 */
	lex.Inc(1)
	if lex.IsEOF() {
		return lex.Errorf("In the middle of a string (an escape sequence no less), and we are EOF. That's not cool")
	}

	if lex.CurrentCharacter() == "x" {
		lex.Inc(3)

		if lex.IsEOF() {
			return lex.Errorf("In the middle of a string (an hex sequence no less), and we are EOF. That's not cool")
		}

		lex.Emit(TOKEN_HEX_CODE)

		/*
		 * Peek ahead a bit
		 */
		if lex.PeekCharacters(1) == "\"" {
			return LexerEndQuote
		}

		return LexerLiteral
	}

	lex.Inc(1)
	if lex.IsEOF() {
		return lex.Errorf("In the middle of a string (an escape sequence no less), and we are EOF. That's not cool")
	}

	lex.Emit(TOKEN_ESCAPED_CHARACTER)
	return LexerLiteral
}
