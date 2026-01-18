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

func (s *Scanner) Tokenize() {
	for !s.isAtEnd() {
		ch := s.peek()

		if isWhiteSpace(ch) { 
			s.pos++
			continue 
		}		
		if ch == byte('{') {
			fmt.Println(tokens.TOKEN_OPEN_BRACE)
		} else if ch == byte('}') {
			fmt.Println(tokens.TOKEN_CLOSE_BRACE)
		} else if ch == byte('[') {
			fmt.Println(tokens.TOKEN_OPEN_BRACKET)
		} else if ch == byte(']') {
			fmt.Println(tokens.TOKEN_CLOSE_BRACKET)
		} else if ch == byte(':') {
			fmt.Println(tokens.TOKEN_COLON)
		} else if ch == byte(',') {
			fmt.Println(tokens.TOKEN_COMMA)
		}
		s.pos++
	}
}

func isWhiteSpace(ch byte) bool {
	SPACE := byte(' ')
	NEW_LINE := byte('\n')
	TAB_SPACE := byte('\t')
	return (ch == SPACE) || (ch == NEW_LINE) || (ch == TAB_SPACE)
}

func (s *Scanner) peek() byte {
	if s.isAtEnd() { return byte('0') }
	return s.data[s.pos]
}

func (s *Scanner) isAtEnd() bool {
	return s.pos >= len(s.data)
}

// ************** Methods End *************** /