package main

import "github.com/adampresley/lexer"

func LexerStartQuote(lex *lexer.Lexer) lexer.LexFn {
	_ = "breakpoint" // LexerStartQuote
	lex.SkipWhitespace()

	if lex.PeekCharacters(1) != "\"" {
		return lex.Errorf("Expected a starting quote... didn't get that.")
	}

	lex.Inc(1)
	if lex.IsEOF() {
		return lex.Errorf("Expected literal or escape sequence. Instead we got EOF'd")
	}

	lex.Emit(TOKEN_QUOTE)

	/*
	 * Next up should either be a literal or an escape sequence. It can also be a blank string
	 */
	if lex.PeekCharacters(1) == "\"" {
		return LexerEndQuote
	}

	if lex.PeekCharacters(1) == "\\" {
		return LexerEscape
	}

	return LexerLiteral
}
