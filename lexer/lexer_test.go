package lexer

import (
	"testing"
    "ksm/token"
)

func TestNew(t *testing.T) {
	tt := []struct {
		input string
		want  *Lexer
	}{
		{"Hello", &Lexer{"Hello", 0, 1, 'H'}},
		{"Godwin", &Lexer{"Godwin", 0, 1, 'G'}},
	}
	for i, tc := range tt {
		got := New(tc.input)
		if got.input != tc.want.input || got.position != tc.want.position || got.readPosition != tc.want.readPosition || got.ch != tc.want.ch {
			t.Fatalf("tests[%d] - Failed: got= %v, want= %v", i, got, tc.want)
		}

	}

}

func TestReadChar(t *testing.T) {
	input := "Godwin"

	l := New(input)

	expectedChars := []struct {
		expectedCh      byte
		expectedPos     int
		expectedReadPos int
	}{
		{'G', 0, 1},
		{'o', 1, 2},
		{'d', 2, 3},
		{'w', 3, 4},
		{'i', 4, 5},
		{'n', 5, 6},
		{0, 6, 7}, // End of input, l.ch should be 0
	}

	for i, expected := range expectedChars {
		if l.ch != expected.expectedCh {

			t.Fatalf("test[%d] - wrong char. expected= %q, got=%q", i, expected.expectedCh, l.ch)
		}
		if l.position != expected.expectedPos {
			t.Fatalf("test[%d] - wrong position, expected= %d, got= %d", i, expected.expectedPos, l.position)
		}
		if l.readPosition != expected.expectedReadPos {
			t.Fatalf("test[%d] - wrong readPosition, expected=  %d, got= %d,", i, expected.expectedReadPos, l.readPosition)
		}
		l.readChar()
	}
}

func TestNextToken(t *testing.T) {
    input := "=+(){},;"

    tests := []struct {
        expectedType token.TokenType
        expectedLiteral string
    }{
        {token.ASSIGN, "="},
        {token.PLUS,"+"},
        {token.LPAREN, "("},
        {token.RPAREN,")"},
        {token.LBRACE,"{"},
        {token.RBRACE,"}"},
        {token.COMMA,","},
        {token.SEMICOLON,";"},
    }
    l := New(input)

    for i, tc := range tests {
        tok := l.NextToken()

        if tok.Type != tc.expectedType {
            t.Fatalf("test[%d] - wrong tokentype. Expected=%q, got=%q", i, tc.expectedType, tok.Type)
        }

        if tok.Literal != tc.expectedLiteral {
            t.Fatalf("test[%d] - wrong literal. Expected %q, got=%q", i, tc.expectedLiteral, tok.Literal)
        }
    }
}

// func TestReadNumber(t *testing.T){
//  	input:= `5;`
// 	tests:=[]struct{
// 		expectedType token.TokenType
// 		expectedLiteral string
// 	}{
// 		{token.INT,"5"},
// 	}

// 	got:=readNumber(input)
// 	for _,tt:=range tests{
// 		if got!= tt.expectedLiteral{
// 			t.Errorf("got: \n want : \n",got,tt.want)
// 		}
// 	}

//  }
