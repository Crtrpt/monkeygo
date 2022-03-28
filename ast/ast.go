package ast

import (
	"monkeygo/token"
)

type Node interface {
	TokenLiteral() string
}

//语句
type Statement interface {
	Node
	statementNode()
}

//表达式
type Expression interface {
	Node
	expressionNode()
}

//一个程序由多个语句组成
type Program struct {
	Statements []Statement
}

//程序
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
