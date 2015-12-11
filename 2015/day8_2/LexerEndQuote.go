package main

import "github.com/adampresley/lexer"

func LexerEndQuote(lex *lexer.Lexer) lexer.LexFn {
	_ = "breakpoint" // LexerEndQuote
	lex.SkipWhitespace()

	if lex.CurrentCharacter() != "\"" {
		return lex.Errorf("Expected an ending quote... didn't get that.")
	}

	lex.Inc(1)
	lex.Emit(TOKEN_QUOTE)
	lex.Discard(1)

	return LexerBegin
}
