package main

import (
	"github.com/adampresley/lexer"
)

const (
	TOKEN_QUOTE lexer.TokenType = iota
	TOKEN_LITERAL
	TOKEN_ESCAPED_CHARACTER
	TOKEN_HEX_CODE
	TOKEN_NEWLINE
)
