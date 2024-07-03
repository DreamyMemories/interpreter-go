package lexer

import (
	"github.com/DreamyMemories/interpreter-go/token"
)

type Lexer struct {
	input        string
	position     int  // current position in input
	readPosition int  // current reading position to input
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// Give next character and advance reader pointer
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var currentToken token.Token
	l.skipWhitespace()
	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			currentToken = token.Token{Type: token.EQ, Literal: string(ch) + string(l.ch)}
		} else {
			currentToken = newToken(token.ASSIGN, l.ch)
		}
	case ';':
		currentToken = newToken(token.SEMICOLON, l.ch)
	case '(':
		currentToken = newToken(token.LPAREN, l.ch)
	case ')':
		currentToken = newToken(token.RPAREN, l.ch)
	case '{':
		currentToken = newToken(token.LBRACE, l.ch)
	case '}':
		currentToken = newToken(token.RBRACE, l.ch)
	case ',':
		currentToken = newToken(token.COMMA, l.ch)
	case '+':
		currentToken = newToken(token.PLUS, l.ch)
	case '-':
		currentToken = newToken(token.MINUS, l.ch)
	case '*':
		currentToken = newToken(token.ASTERISK, l.ch)
	case '/':
		currentToken = newToken(token.SLASH, l.ch)
	case '<':
		currentToken = newToken(token.LT, l.ch)
	case '>':
		currentToken = newToken(token.GT, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			currentToken = token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(l.ch)}
		} else {
			currentToken = newToken(token.BANG, l.ch)
		}
	case 0:
		currentToken.Literal = ""
		currentToken.Type = token.EOF
	default:
		if isLetter(l.ch) {
			currentToken.Literal = l.readIdentifier()
			currentToken.Type = token.LookupIdent(currentToken.Literal)
			return currentToken
		} else if isDigit(l.ch) {
			currentToken.Type = token.INT
			currentToken.Literal = l.readNumber()
			return currentToken
		} else {
			currentToken = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return currentToken
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\n' || l.ch == '\t' || l.ch == '\r' {
		l.readChar()
	}
}

func isLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z' || ('A' <= ch && 'Z' <= ch) || ch == '_')
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Literal: string(ch), Type: tokenType}
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}
