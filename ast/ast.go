package ast

import (
	"bytes"
	"djinn/token"
)

//Every Node in our AST implemenets the Node interface
type Node interface {
	TokenLiteral() string
	String() string
}

//Some Nodes can be Statements
type Statement interface {
	Node
	statementNode()
}

//Some Nodes can be Expressions
type Expression interface {
	Node
	expressionNode()
}

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

type Identifier struct {
	Token token.Token // The token.IDENT Token
	Value string
}

type PrefixExpression struct {
	Token    token.Token `prefix token ex. !`
	Operator string      `json:"operator,omitempty"`
	Right    Expression  `The expression to the right of the current extension`
}

type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator string
	Right    Expression
}

type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

type CreateStatement struct {
	Token token.Token // The token.Create Token
	Name  *Identifier
	Value Expression
}

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

//Program Node is going to be the root node of every AST our parser produces
type Program struct {
	Statements []Statement
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")
	return out.String()
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

func (cs *CreateStatement) String() string {
	var out bytes.Buffer
	out.WriteString(cs.TokenLiteral() + " ")
	out.WriteString(cs.Name.String())
	out.WriteString(" = ")
	if cs.Value != nil {
		out.WriteString(cs.Value.String())
	}
	out.WriteString(";")
	return out.String()
}

func (cs *CreateStatement) statementNode()       {}
func (cs *CreateStatement) TokenLiteral() string { return cs.Token.Literal }

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

func (i *Identifier) String() string { return i.Value }

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string       { return il.Token.Literal }

func (pe *PrefixExpression) TokenLiteral() string {
	return pe.Token.Literal
}

func (pe *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

func (pe *PrefixExpression) expressionNode() {

}
func (ie *InfixExpression) TokenLiteral() string {
	return ie.Token.Literal
}

func (ie *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")

	return out.String()
}

func (ie *InfixExpression) expressionNode() {}
