package main

import "github.com/adampresley/lexer"

func LexerBegin(lex *lexer.Lexer) lexer.LexFn {
	_ = "breakpoint" // LexerBegin
	lex.SkipWhitespace()

	if lex.IsEOF() {
		return LexerEnd
	}

	return LexerStartQuote
}
