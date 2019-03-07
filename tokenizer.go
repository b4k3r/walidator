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

var keywords = [...]string{
	"SubClassOf", "EquivalentClasses", "DisjointClasses", "SameIndividual", "DifferentIndividuals",
	"ObjectIntersectionOf", "ObjectUnionOf", "ObjectComplementOf", "ObjectOneOf"}
var isAlpha = regexp.MustCompile("^[a-zA-Z]+").MatchString
var isSeparator = regexp.MustCompile("[(|)]").MatchString

func (t *Tokenizer) Scan() {
	var s scanner.Scanner
	t.currentLine = 1
	t.currentColumn = 0

	s.Init(bytes.NewReader(t.content))
	s.Filename = t.fileName

	for tok := s.Next(); tok != scanner.EOF; tok = s.Next() {
		t.currentColumn++

		if isAlpha(string(tok)) {
			literalInRunes := []rune{tok}

			for {
				tok = s.Peek()
				if isAlpha(string(tok)) {
					literalInRunes = append(literalInRunes, tok)
					s.Next()
					t.currentColumn++
				} else {
					break
				}
			}

			literal := string(literalInRunes)

			if isKeyword(literal) {
				t.addToken("KEYWORD", literal)
			} else {
				t.addToken("ALPHA", literal)
			}
		} else if tok == ' ' {
			continue
		} else if tok == '\n' {
			t.currentColumn = 0
			t.currentLine++
			continue
		} else if tok == ':' {
			t.addToken(":", ":")
		} else if isSeparator(string(tok)) {
			t.addToken(string(tok), string(tok))
		} else {
			fmt.Fprintf(os.Stderr, "%s:%d:%d unexpected charater '%s' \n", t.fileName, t.currentLine, t.currentColumn, string(tok))
			os.Exit(1)
		}
	}
	t.addToken("EOF", "EOF")
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
