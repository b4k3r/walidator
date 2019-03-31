package main

import (
	"fmt"
	"os"
)

type Parser struct {
	tokens     []Token
	token      Token
	tokenIndex int
	tokenizer  Tokenizer
}

func NewParser(t Tokenizer) *Parser {
	p := new(Parser)
	p.tokens = t.tokens
	p.tokenIndex = 0
	p.token = p.tokens[0]
	p.tokenizer = t

	return p
}

var axiomKeywords = [...]string{
	"SubClassOf", "EquivalentClasses", "DisjointClasses", "SameIndividual", "DifferentIndividuals"}
var classExpressionsKeywords = [...]string{
	"ObjectIntersectionOf", "ObjectUnionOf", "ObjectComplementOf", "ObjectOneOf"}

func (p *Parser) start() {
	p.program()

	fmt.Println("Syntax OK!")
}

func (p *Parser) program() {
	p.axiom()

	if p.isAxiomKeyword() {
		p.program()
	} else if p.token.value != "EOF" {
		p.syntaxError("<axiom>")
	}
}

func (p *Parser) axiom() {
	if p.token.value == "SubClassOf" {
		p.subClassOf()
	} else if p.token.value == "EquivalentClasses" {
		p.equivalentClasses()
	} else if p.token.value == "DisjointClasses" {
		p.disjointClasses()
	} else if p.token.value == "SameIndividual" {
		p.sameIndividual()
	} else if p.token.value == "DifferentIndividuals" {
		p.differentIndividuals()
	}
}

func (p *Parser) subClassOf() {
	p.takeToken("KEYWORD")
	p.takeToken("(")
	p.classExpression()
	p.classExpression()
	p.takeToken(")")
}

func (p *Parser) equivalentClasses() {
	p.takeToken("KEYWORD")
	p.takeToken("(")
	p.classExpression()
	p.classExpression()
	p.classExpressions()
	p.takeToken(")")
}

func (p *Parser) disjointClasses() {
	p.takeToken("KEYWORD")
	p.takeToken("(")
	p.classExpression()
	p.classExpression()
	p.classExpressions()
	p.takeToken(")")
}

func (p *Parser) sameIndividual() {
	p.takeToken("KEYWORD")
	p.takeToken("(")
	p.id()
	p.id()
	p.ids()
	p.takeToken(")")
}

func (p *Parser) differentIndividuals() {
	p.takeToken("KEYWORD")
	p.takeToken("(")
	p.id()
	p.id()
	p.ids()
	p.takeToken(")")
}

func (p *Parser) id() {
	p.takeToken(":")
	p.takeToken("ALPHA")
}

func (p *Parser) ids() {
	if p.token.value == ":" {
		p.id()
		p.ids()
	}
}

func (p *Parser) classExpression() {
	if p.token.value == ":" {
		p.id()
	} else if p.token.value == "ObjectIntersectionOf" {
		p.objectIntersectionOf()
	} else if p.token.value == "ObjectUnionOf" {
		p.objectUnionOf()
	} else if p.token.value == "ObjectComplementOf" {
		p.objectComplementOf()
	} else if p.token.value == "ObjectOneOf" {
		p.objectOneOf()
	} else {
		p.syntaxError("<classExpression>")
	}
}

func (p *Parser) classExpressions() {
	if p.token.value == ":" || p.isClassExpressionsKeyword() {
		p.classExpression()
		p.classExpressions()
	}
}

func (p *Parser) objectIntersectionOf() {
	p.takeToken("KEYWORD")
	p.takeToken("(")
	p.classExpression()
	p.classExpression()
	p.classExpressions()
	p.takeToken(")")
}

func (p *Parser) objectUnionOf() {
	p.takeToken("KEYWORD")
	p.takeToken("(")
	p.classExpression()
	p.classExpression()
	p.classExpressions()
	p.takeToken(")")
}

func (p *Parser) objectComplementOf() {
	p.takeToken("KEYWORD")
	p.takeToken("(")
	p.classExpression()
	p.takeToken(")")
}

func (p *Parser) objectOneOf() {
	p.takeToken("KEYWORD")
	p.takeToken("(")
	p.id()
	p.ids()
	p.takeToken(")")
}

func (p *Parser) takeToken(tokenType string) {
	if p.token.kind != tokenType {
		p.syntaxError(tokenType)
	}

	p.tokenIndex++
	p.token = p.tokens[p.tokenIndex]
}

func (p *Parser) syntaxError(expectedToken string) {
	fmt.Fprintf(os.Stderr, "%s:%d:%d: syntax error, unexpected symbol '%s', expected: '%s' symbol \n",
		p.tokenizer.fileName, p.token.line, p.token.column, p.token.value, expectedToken)
	os.Exit(1)
}

func (p *Parser) peekToken() Token {
	return p.tokens[p.tokenIndex+1]
}

func (p *Parser) isAxiomKeyword() bool {
	for _, ak := range axiomKeywords {
		if ak == p.token.value {
			return true
		}
	}
	return false
}

func (p *Parser) isClassExpressionsKeyword() bool {
	for _, ak := range classExpressionsKeywords {
		if ak == p.token.value {
			return true
		}
	}
	return false
}
