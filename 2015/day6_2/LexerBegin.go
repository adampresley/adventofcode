package main

import (
	"github.com/adampresley/lexer"
)

func LexerBegin(lex *lexer.Lexer) lexer.LexFn {
	_ = "breakpoint"
	lex.SkipWhitespace()

	if lex.IsEOF() {
		return nil
	}

	return LexerCommand
}
