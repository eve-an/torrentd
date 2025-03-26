package lexer

import (
	"testing"

	"github.com/eve-an/torrentd/pkg/token"
	"github.com/stretchr/testify/assert"
)

func TestLexer_Lex_Integers(t *testing.T) {
	tests := []struct {
		name string
		data string
		want []token.Token
	}{
		{
			"simple integer",
			"i32e",
			[]token.Token{
				{Value: "i", Type: token.IntegerStart},
				{Value: "32", Type: token.Integer},
				{Value: "e", Type: token.IntegerEnd},
			},
		},
		{
			"simple negative integer",
			"i-32e",
			[]token.Token{
				{Value: "i", Type: token.IntegerStart},
				{Value: "-32", Type: token.Integer},
				{Value: "e", Type: token.IntegerEnd},
			},
		},
		{
			"simple zero integer",
			"i0e",
			[]token.Token{
				{Value: "i", Type: token.IntegerStart},
				{Value: "0", Type: token.Integer},
				{Value: "e", Type: token.IntegerEnd},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewLexer()
			got, err := l.Lex(tt.data)

			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestLexer_Lex_Strings(t *testing.T) {
	tests := []struct {
		name string
		data string
		want []token.Token
	}{
		{
			"simple string",
			"3:SOS",
			[]token.Token{
				{Value: "3", Type: token.StringLength},
				{Value: ":", Type: token.StringSeperator},
				{Value: "SOS", Type: token.String},
			},
		},
		{
			"empty string",
			"0:",
			[]token.Token{
				{Value: "0", Type: token.StringLength},
				{Value: ":", Type: token.StringSeperator},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewLexer()
			got, err := l.Lex(tt.data)

			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestLexer_Lex_List(t *testing.T) {
	tests := []struct {
		name string
		data string
		want []token.Token
	}{
		{
			"simple list",
			"l4:ABBAi-20ee",
			[]token.Token{
				{Value: "l", Type: token.ListStart},
				{Value: "4", Type: token.StringLength},
				{Value: ":", Type: token.StringSeperator},
				{Value: "ABBA", Type: token.String},
				{Value: "i", Type: token.IntegerStart},
				{Value: "-20", Type: token.Integer},
				{Value: "e", Type: token.IntegerEnd},
				{Value: "e", Type: token.ListEnd},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewLexer()
			got, err := l.Lex(tt.data)

			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestLexer_Lex_Dicionary(t *testing.T) {
	tests := []struct {
		name string
		data string
		want []token.Token
	}{
		{
			"simple dictionary",
			"d4:ABBAi-20ee",
			[]token.Token{
				{Value: "d", Type: token.DictionaryStart},
				{Value: "4", Type: token.StringLength},
				{Value: ":", Type: token.StringSeperator},
				{Value: "ABBA", Type: token.String},
				{Value: "i", Type: token.IntegerStart},
				{Value: "-20", Type: token.Integer},
				{Value: "e", Type: token.IntegerEnd},
				{Value: "e", Type: token.DictionaryEnd},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewLexer()
			got, err := l.Lex(tt.data)

			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
