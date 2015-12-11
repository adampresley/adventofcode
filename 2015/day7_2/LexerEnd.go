package main

import (
	"github.com/adampresley/lexer"
)

/*
LexerEnd means we've reached the end of the line. Emit an EOF and
return nil. This tells the lexer to stop the goroutine that is lexing.
*/
func LexerEnd(lex *lexer.Lexer) lexer.LexFn {
	lex.Emit(lexer.TOKEN_EOF)
	return nil
}
