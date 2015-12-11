package main

import "github.com/adampresley/lexer"

func LexerLiteral(lex *lexer.Lexer) lexer.LexFn {
	_ = "breakpoint" // LexerLiteral

	for {
		if lex.IsEOF() {
			return lex.Errorf("In the middle of a string, and we are EOF. That's not cool")
		}

		if lex.CurrentCharacter() == "\\" {
			lex.Emit(TOKEN_LITERAL)
			return LexerEscape
		}

		if lex.CurrentCharacter() == "\"" {
			lex.Emit(TOKEN_LITERAL)
			return LexerEndQuote
		}

		lex.Inc(1)
	}
}
