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

	start := s.pos
	ch := s.consume()
	end := s.pos

	switch ch {
	case '{':
		return tokens.NewToken(tokens.TOKEN_OPEN_BRACE, start, end)
	case '}':
		return tokens.NewToken(tokens.TOKEN_CLOSE_BRACE, start, end)
	case '[':
		return tokens.NewToken(tokens.TOKEN_OPEN_BRACKET, start, end)
	case ']':
		return tokens.NewToken(tokens.TOKEN_CLOSE_BRACKET, start, end)
	case ':':
		return tokens.NewToken(tokens.TOKEN_COLON, start, end)
	case ',':
		return tokens.NewToken(tokens.TOKEN_COMMA, start, end)
	case '"':
		// handle string (need to go through the characters until you reach closing string)
		for !s.isAtEnd() && s.peek() != '"' {
			s.advance()
		}
		return tokens.NewToken(tokens.TOKEN_STRING, start, end)
	}
	panic(fmt.Sprintf("Unexpected character '%c' at position %d", ch, s.pos-1))
}

func isWhiteSpace(ch byte) bool {
	SPACE := byte(' ')
	NEW_LINE := byte('\n')
	TAB_SPACE := byte('\t')
	return (ch == SPACE) || (ch == NEW_LINE) || (ch == TAB_SPACE)
}

func (s *Scanner) advance() byte {
	curr_byte := s.data[s.pos]
	s.pos++
	return curr_byte
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