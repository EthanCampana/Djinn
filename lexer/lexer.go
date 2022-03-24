package lexer

import (
	"djinn/token"
)

type Lexer struct {
	input        string // Input Source Code that is getting interpreted
	position     int    // Current position of input.(This input has already been read)
	readPosition int    // Current reading postion(On After the position) this is the char that is going to be read
	char         byte   // Current char in bytes that is currently going to be examined
}

func New(in string) *Lexer {
	l := &Lexer{input: in}
	l.readChar()
	return l
}

//Reads the char at the reading postions and saves it to the char byte. If reading position
//is beyond the max length of the input we set char to NULL because we reached EOF or EOL
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.char = 0
	} else {
		l.char = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func newToken(t token.TokenType, char byte) token.Token {
	return token.Token{Type: t, Literal: string(char)}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhitespace()
	switch l.char {

	default:
		if isLetter(l.char) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.char) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.char)
		}
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()
	case '=':
		if l.peakChar() == '=' {
			temp_tok := l.char
			l.readChar()
			literal := string(temp_tok) + string(l.char)
			tok = token.Token{Type: token.EQ, Literal: literal}
		} else {
			tok = newToken(token.ASSIGN, l.char)
		}

	case '+':
		tok = newToken(token.PLUS, l.char)
	case ';':
		tok = newToken(token.SEMICOLON, l.char)
	case ',':
		tok = newToken(token.COMMA, l.char)
	case '(':
		tok = newToken(token.LPAREN, l.char)
	case ')':
		tok = newToken(token.RPAREN, l.char)
	case '{':
		tok = newToken(token.LBRACE, l.char)
	case '}':
		tok = newToken(token.RBRACE, l.char)
	case '<':
		tok = newToken(token.LT, l.char)
	case '>':
		tok = newToken(token.GT, l.char)
	case '!':
		if l.peakChar() == '=' {
			temp_tok := l.char
			l.readChar()
			literal := string(temp_tok) + string(l.char)
			tok = token.Token{Type: token.EQ, Literal: literal}
		} else {
			tok = newToken(token.BANG, l.char)
		}
	case '*':
		tok = newToken(token.ASTERISK, l.char)
	case '-':
		tok = newToken(token.MINUS, l.char)
	case '/':
		tok = newToken(token.SLASH, l.char)
	case 0:
		tok = newToken(token.EOF, l.char)
	case '[':
		tok = newToken(token.LBRACKET, l.char)
	case ']':
		tok = newToken(token.RBRACKET, l.char)
	case ':':
		tok = newToken(token.COLON, l.char)
	}

	//After we get the token we want to increment the read position to the next token to read
	l.readChar()
	return tok

}

// Builds the String to put it into the token literal
func (l *Lexer) readString() string {

	position := l.position + 1
	for {
		l.readChar()
		if l.char == '"' || l.char == 0 {
			break
		}
	}
	return l.input[position:l.position]

}

func (l *Lexer) readIdentifier() string {
	// Create a new variable that will hold the position of the first char in the keyword
	position := l.position
	//Read the rest of the letters in the word and increment the l.position till the end of the word
	for isLetter(l.char) {
		l.readChar()
	}
	//return the splice of the start postion of the first char till the end of the word
	return l.input[position:l.position]

}

func isLetter(char byte) bool {
	//Check if the char is a letter doing some byte math b(0-0=)b
	return 'a' <= char && 'z' >= char || 'A' <= char && 'Z' >= char || char == '_'
}

func (l *Lexer) skipWhitespace() {
	for l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readNumber() string {
	// Create a new variable that will hold the position of the first char in the keyword
	position := l.position
	//Read the rest of the letters in the word and increment the l.position till the end of the word
	for isDigit(l.char) {
		l.readChar()
	}
	//return the splice of the start postion of the first char till the end of the word
	return l.input[position:l.position]

}

func isDigit(char byte) bool {
	return '0' <= char && char <= '9'
}

//Use this function to peak at the next char in the input of the lexer
func (l *Lexer) peakChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}
