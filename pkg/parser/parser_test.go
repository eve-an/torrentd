package parser_test

import (
	"testing"

	"github.com/eve-an/torrentd/pkg/bencoding"
	"github.com/eve-an/torrentd/pkg/parser"
	"github.com/eve-an/torrentd/pkg/token"
	"github.com/stretchr/testify/assert"
)

func TestParser_Parse(t *testing.T) {
	tests := []struct {
		name   string
		tokens []token.Token
		want   bencoding.Dict
	}{
		{
			name: "simple",
			tokens: []token.Token{
				{Type: token.DictionaryStart, Value: "d"},
				{Type: token.StringLength, Value: "4"},
				{Type: token.StringSeperator, Value: ":"},
				{Type: token.String, Value: "info"},
				{Type: token.IntegerStart, Value: "i"},
				{Type: token.Integer, Value: "40"},
				{Type: token.IntegerEnd, Value: "e"},
				{Type: token.StringLength, Value: "4"},
				{Type: token.StringSeperator, Value: ":"},
				{Type: token.String, Value: "abba"},
				{Type: token.StringLength, Value: "4"},
				{Type: token.StringSeperator, Value: ":"},
				{Type: token.String, Value: "abcd"},
				{Type: token.DictionaryEnd, Value: "e"},
			},
			want: bencoding.Dict{
				Entities: map[string]bencoding.Value{
					"info": bencoding.Integer{Value: 40},
					"abba": bencoding.String{Value: "abcd"},
				},
			},
		},
		{
			name: "dict in dict",
			tokens: []token.Token{
				{Type: token.DictionaryStart, Value: "d"},
				{Type: token.StringLength, Value: "4"},
				{Type: token.StringSeperator, Value: ":"},
				{Type: token.String, Value: "info"},
				{Type: token.DictionaryStart, Value: "d"},
				{Type: token.StringLength, Value: "4"},
				{Type: token.StringSeperator, Value: ":"},
				{Type: token.String, Value: "abcd"},
				{Type: token.StringLength, Value: "4"},
				{Type: token.StringSeperator, Value: ":"},
				{Type: token.String, Value: "efgh"},
				{Type: token.DictionaryEnd, Value: "e"},
				{Type: token.DictionaryEnd, Value: "e"},
			},
			want: bencoding.Dict{
				Entities: map[string]bencoding.Value{
					"info": bencoding.Dict{Entities: map[string]bencoding.Value{
						"abcd": bencoding.String{Value: "efgh"},
					}},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := parser.NewParser(tt.tokens)

			value, err := p.Parse()

			assert.NoError(t, err)
			assert.Equal(t, tt.want, value)
		})
	}
}
