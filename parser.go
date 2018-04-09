package main

import (
	"fmt"
	"os"
)

type Parser struct {
	tokens     []Token
	token      Token
	tokenIndex int
}

func NewParser(tokens []Token) *Parser {
	p := new(Parser)
	p.tokens = tokens
	p.tokenIndex = 0
	p.token = p.tokens[0]

	return p
}

func (p *Parser) start() {
	p.program()
	p.takeToken("SEMICOLLON")
	fmt.Println("start OK!")
}

func (p *Parser) program() {
	p.dec()
	p.takeToken("EQ")
	p.fromClause()
	p.whereClause()
	p.orderbyClause()
	p.selectClause()
	fmt.Println("program OK!")
}

func (p *Parser) dec() {
	p.takeToken("DEC")
	p.takeToken("ID")

	fmt.Println("dec OK!")
}

func (p *Parser) fromClause() {
	if p.token.kind == "from" {
		p.takeToken("from")
		p.takeToken("ID")
		p.takeToken("in")
		p.arg()

		fmt.Println("from_clause OK!")
	} else {
		fmt.Fprintln(os.Stderr, "Epsilon is not allowed")
	}
}

func (p *Parser) arg() {
	p.takeToken("ID")

	if p.token.kind == "DOT" {
		p.takeToken("DOT")

		if p.peekToken().kind == "(" {
			p.method()
		} else {
			p.arg()
		}
	}
}

func (p *Parser) whereClause() {
	if p.token.kind == "where" {
		p.takeToken("where")

		if p.peekToken().kind == "OPERATOR" {
			p.arg()
			p.takeToken("OPERATOR")
			p.arg()
		} else if p.token.kind == "NEGATION" {
			p.negation()
			p.arg()
		} else {
			p.negation()
			p.arg()
		}

		fmt.Println("where_clause OK!")
	} else {
		fmt.Println("where_clause skipped!")
	}
}

func (p *Parser) negation() {
	if p.token.kind == "NEGATION" {
		p.takeToken("NEGATION")
	}
}

func (p *Parser) method() {
	p.takeToken("ID")
	p.takeToken("(")
	p.arg()
	p.takeToken(")")
}

func (p *Parser) orderbyClause() {
	if p.token.kind == "orderby" {
		p.takeToken("orderby")

		p.arg()
		p.orderType()

		fmt.Println("orderby_clause OK!")
	} else {
		fmt.Println("orderby_clause skipped!")
	}
}

func (p *Parser) orderType() {
	if p.token.kind == "descending" {
		p.takeToken("descending")
	} else if p.token.kind == "ascending" {
		p.takeToken("ascending")
	} else {
		fmt.Println("orderby_type skipped!")
	}
}

func (p *Parser) selectClause() {
	p.takeToken("select")
	p.selectArgs()
	p.selectMethod()

	fmt.Println("select_clause OK!")
}

func (p *Parser) selectMethod() {
	if p.token.kind == "DOT" {
		p.takeToken("DOT")
		p.takeToken("ID")
		p.takeToken("(")
		p.takeToken(")")
	} else {
		fmt.Println("select_method skipped!")
	}

}

func (p *Parser) selectArgs() {
	if p.token.kind == "new" {
		p.takeToken("new")
		p.takeToken("{")
		p.selectNewArg()
		p.takeToken("}")
	} else {
		p.arg()
	}
}

func (p *Parser) selectNewArg() {
	p.arg()
	p.selectNewArgs()
}

func (p *Parser) selectNewArgs() {
	if p.token.kind == "COMMA" {
		p.takeToken("COMMA")
		p.arg()
		p.selectNewArgs()
	} else if p.token.kind == "EQ" {
		p.takeToken("EQ")
		p.arg()
		p.selectNewArgs()
	}
}

func (p *Parser) takeToken(tokenType string) {
	if p.token.kind != tokenType {
		fmt.Fprintf(os.Stderr, "Unexpected token '%s', line: %d, column: %d \n", p.token.kind, p.token.line, p.token.column)
		os.Exit(1)
	}

	p.token = p.nextToken()
}

func (p *Parser) nextToken() Token {
	if p.tokenIndex < len(p.tokens)-1 {
		p.tokenIndex++
	}

	return p.tokens[p.tokenIndex]
}

func (p *Parser) peekToken() Token {
	return p.tokens[p.tokenIndex+1]
}
