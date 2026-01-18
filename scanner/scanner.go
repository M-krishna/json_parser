package scanner

import (
	"fmt"
	"json_parser/tokens"
)

type Scanner struct {
	data []byte
	pos int
}

func NewScanner(data []byte) Scanner {
	return Scanner{data: data}
}

// ************** Methods Start *************** /

func (s *Scanner) NextToken() tokens.Token {
	for !s.isAtEnd() && isWhiteSpace(s.peek()) {
		s.advance()
	}

	if s.isAtEnd() {
		return tokens.NewToken(tokens.TOKEN_EOF, s.pos, s.pos)
	}

	ch := s.consume()

	switch ch {
	case '{':
		return tokens.NewToken(tokens.TOKEN_OPEN_BRACE, 0, 0)
	case '}':
		return tokens.NewToken(tokens.TOKEN_CLOSE_BRACE, 0, 0)
	case '[':
		return tokens.NewToken(tokens.TOKEN_OPEN_BRACKET, 0, 0)
	case ']':
		return tokens.NewToken(tokens.TOKEN_CLOSE_BRACKET, 0, 0)
	case ':':
		return tokens.NewToken(tokens.TOKEN_COLON, 0, 0)
	case ',':
		return tokens.NewToken(tokens.TOKEN_COMMA, 0, 0)
	}
	panic(fmt.Sprintf("Unexpected character '%c' at position %d", ch, s.pos-1))
}

func isWhiteSpace(ch byte) bool {
	SPACE := byte(' ')
	NEW_LINE := byte('\n')
	TAB_SPACE := byte('\t')
	return (ch == SPACE) || (ch == NEW_LINE) || (ch == TAB_SPACE)
}

func (s *Scanner) advance() {
	s.pos++
}

func (s *Scanner) consume() byte {
	ch := s.peek()
	s.advance()
	return ch
}

func (s *Scanner) peek() byte {
	if s.isAtEnd() { return byte('0') }
	return s.data[s.pos]
}

func (s *Scanner) isAtEnd() bool {
	return s.pos >= len(s.data)
}

// ************** Methods End *************** /