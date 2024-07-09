package parser

import (
	"testing"

	"github.com/DreamyMemories/interpreter-go/ast"
	"github.com/DreamyMemories/interpreter-go/lexer"
)

func TestLetStatement(t *testing.T) {
	input := `
		let x = 5;
		let y = 10;
		let foobar = 123;
	`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()

	if program == nil {
		t.Fatalf("ParseProgram returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("Program statement does not contain 3 statements. Got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		statement := program.Statements[i]
		if !testLetStatement(t, statement, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let' but got=%q instead", s.TokenLiteral())
		return false
	}

	letStatement, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement but got=%q instead", s)
		return false
	}

	if letStatement.Name.Value != name {
		t.Errorf("s.name not '%s' but instead got='%s'", name, letStatement.Name)
		return false
	}

	if letStatement.Name.TokenLiteral() != name {
		t.Errorf("s.Name not '%s'. got = %s", name, letStatement.Name)
		return false
	}

	return true
}
