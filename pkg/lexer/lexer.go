package lexer

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/eve-an/torrentd/pkg/collection"
	"github.com/eve-an/torrentd/pkg/token"
)

type Lexer struct {
	stack *collection.Stack[token.Type]
}

func NewLexer() *Lexer {
	return &Lexer{
		stack: collection.NewStack[token.Type](128),
	}
}

func (l *Lexer) Lex(data string) ([]token.Token, error) {
	var t []token.Token
	var err error

	tokens := make([]token.Token, 0, 512)

	for len(data) > 0 {
		switch data[0] {
		case 'i':
			t, data, err = l.lexInteger(data)
		case 'l':
			t, data, err = l.lexList(data)
		case 'd':
			t, data, err = l.lexDictionary(data)
		case 'e':
			t, data, err = l.handleEndToken(data)
		default:
			t, data, err = l.lexByteString(data)
		}

		if err != nil {
			return nil, err
		}

		tokens = append(tokens, t...)
	}

	return tokens, nil
}

func (l *Lexer) handleEndToken(data string) ([]token.Token, string, error) {
	tokenType, err := l.stack.Pop()
	if err != nil {
		return nil, "", err
	}

	return []token.Token{{Value: string(data[0]), Type: tokenType}}, data[1:], nil
}

func (l *Lexer) lexByteString(s string) ([]token.Token, string, error) {
	idx := strings.Index(s, ":")

	if idx == -1 {
		return nil, "", fmt.Errorf("could not find separator for byte string '%s'", s)
	}

	rawStringLength := s[:idx]
	stringLength, err := strconv.ParseInt(rawStringLength, 10, 64)
	if err != nil {
		return nil, "", err
	}

	tokens := []token.Token{
		{
			Value: rawStringLength,
			Type:  token.StringLength,
		},
		{
			Value: string(s[idx]),
			Type:  token.StringSeperator,
		},
	}

	if stringLength > 0 {
		tokens = append(tokens, token.Token{
			Value: s[idx+1 : idx+1+int(stringLength)],
			Type:  token.String,
		})
	}

	return tokens, s[idx+1+int(stringLength):], nil
}

func (l *Lexer) lexDictionary(data string) ([]token.Token, string, error) {
	l.stack.Push(token.DictionaryEnd)

	if len(data) < 1 {
		return nil, "", fmt.Errorf("unexpected EOF for dictionary lexing: %s", data)
	}

	return []token.Token{{
		Value: string(data[0]),
		Type:  token.DictionaryStart,
	}}, data[1:], nil
}

func (l *Lexer) lexInteger(data string) ([]token.Token, string, error) {
	idx := strings.Index(data, "e")

	if idx == -1 {
		return nil, "", fmt.Errorf("no end token found for integer: %s", data)
	}

	return []token.Token{
		{
			Value: string(data[0]),
			Type:  token.IntegerStart,
		},
		{
			Value: string(data[1:idx]),
			Type:  token.Integer,
		},
		{
			Value: string(data[idx]),
			Type:  token.IntegerEnd,
		},
	}, data[min(idx+1, len(data)):], nil
}

func (l *Lexer) lexList(data string) ([]token.Token, string, error) {
	l.stack.Push(token.ListEnd)

	if len(data) < 1 {
		return nil, "", fmt.Errorf("unexpected EOF for list lexing: %s", data)
	}

	return []token.Token{{
		Value: string(data[0]),
		Type:  token.ListStart,
	}}, data[1:], nil
}
