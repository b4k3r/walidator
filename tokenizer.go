package main

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"text/scanner"
)

type Token struct {
	kind   string
	value  string
	line   int
	column int
}

type Tokenizer struct {
	content       []byte
	fileName      string
	tokens        []Token
	currentLine   int
	currentColumn int
}

var keywords = [...]string{"where", "in", "from", "descending", "ascending", "orderby", "select", "new"}
var isID = regexp.MustCompile("^[a-zA-Z0-9_]+").MatchString
var isOperator = regexp.MustCompile("!=|>|<|>=|%|<=|=").MatchString
var isSeparator = regexp.MustCompile("[{|}|(|)]").MatchString

func (t *Tokenizer) Scan() {
	var s scanner.Scanner
	t.currentLine = 1
	t.currentColumn = 0

	s.Init(bytes.NewReader(t.content))
	s.Filename = t.fileName

	for tok := s.Next(); tok != scanner.EOF; tok = s.Next() {
		t.currentColumn++

		if isID(string(tok)) {
			literal_in_runes := []rune{tok}

			for {
				tok = s.Peek()
				if isID(string(tok)) {
					literal_in_runes = append(literal_in_runes, tok)
					s.Next()
					t.currentColumn++
				} else {
					break
				}
			}

			literal := string(literal_in_runes)

			if isKeyword(literal) {
				t.addToken(literal, literal)
			} else if string(literal) == "var" {
				t.addToken("DEC", literal)
			} else {
				t.addToken("ID", literal)
			}
		} else if tok == ' ' {
			continue
		} else if tok == '\n' {
			t.currentColumn = 0
			t.currentLine++
			continue
		} else if tok == '=' {
			t.addToken("EQ", "=")
		} else if isOperator(string(tok)) {
			t.addToken("OPERATOR", string(tok))
		} else if tok == ';' {
			t.addToken("SEMICOLLON", "=")
		} else if tok == ',' {
			t.addToken("COMMA", ",")
		} else if tok == '.' {
			t.addToken("DOT", ".")
		} else if tok == '!' {
			t.addToken("NEGATION", "!")
		} else if isSeparator(string(tok)) {
			t.addToken(string(tok), string(tok))
		} else {
			fmt.Fprintf(os.Stderr, "%s:%d:%d Unexpected charater '%s' \n", t.fileName, t.currentLine, t.currentColumn, string(tok))
			os.Exit(1)
		}
	}
}

func isKeyword(str string) bool {
	for _, keyw := range keywords {
		if keyw == str {
			return true
		}
	}
	return false
}

func (t *Tokenizer) addToken(kind string, value string) {
	token := Token{kind: kind, value: value, line: t.currentLine, column: t.currentColumn}

	t.tokens = append(t.tokens, token)
}
