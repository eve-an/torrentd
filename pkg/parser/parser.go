package parser

import (
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/eve-an/torrentd/pkg/bencoding"
	"github.com/eve-an/torrentd/pkg/lexer"
	"github.com/eve-an/torrentd/pkg/token"
)

var EOT = errors.New("EOT") // end of token

type Parser struct {
	tokens []token.Token
	pos    int
}

func NewParser(tokens []token.Token) *Parser {
	return &Parser{
		tokens: tokens,
		pos:    0,
	}
}

func ParseFromReader(r io.Reader) (bencoding.Value, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	tokens, err := lexer.NewLexer().Lex(string(data))
	if err != nil {
		return nil, err
	}

	return NewParser(tokens).Parse()
}

func (p *Parser) expect(t token.Type) (token.Token, error) {
	if p.pos >= len(p.tokens) {
		return token.Token{}, errors.New("unexpected end of tokens")
	}

	current := p.tokens[p.pos]
	if current.Type != t {
		return token.Token{}, fmt.Errorf("expected token %+v, got %+v", t, current.Type)
	}

	p.pos++
	return current, nil
}

func (p *Parser) Parse() (bencoding.Value, error) {
	tok, err := p.current()
	if err != nil {
		return nil, err
	}

	switch tok.Type {
	case token.IntegerStart:
		return p.parseInteger()
	case token.StringLength:
		return p.parseString()
	case token.ListStart:
		return p.parseList()
	case token.DictionaryStart:
		return p.parseDictionary()
	default:
		return nil, fmt.Errorf("unexpected token type: %+v", tok.Type)
	}
}

func (p *Parser) parseDictionary() (bencoding.Value, error) {
	if _, err := p.expect(token.DictionaryStart); err != nil {
		return nil, err
	}

	entities := make(map[string]bencoding.Value, 16)

	for {
		tok, err := p.current()
		if err != nil {
			return nil, err
		}

		if tok.Type == token.DictionaryEnd {
			break
		}

		key, err := p.parseString()
		if err != nil {
			return nil, err
		}

		value, err := p.Parse()
		if err != nil {
			return nil, err
		}

		entities[key.(bencoding.String).Value] = value
	}

	if _, err := p.expect(token.DictionaryEnd); err != nil {
		return nil, err
	}

	return bencoding.Dict{Entities: entities}, nil
}

func (p *Parser) hasTokens() bool {
	return p.pos < len(p.tokens)
}

func (p *Parser) current() (token.Token, error) {
	if !p.hasTokens() {
		return token.Token{}, EOT
	}

	return p.tokens[p.pos], nil
}

func (p *Parser) parseList() (bencoding.Value, error) {
	if _, err := p.expect(token.ListStart); err != nil {
		return nil, err
	}

	items := make([]bencoding.Value, 0, 16)

	for {
		tok, err := p.current()
		if err != nil {
			return nil, err
		}

		if tok.Type == token.ListEnd {
			break
		}

		nextValue, err := p.Parse()
		if err != nil {
			return nil, err
		}

		items = append(items, nextValue)
	}

	if _, err := p.expect(token.ListEnd); err != nil {
		return nil, err
	}

	return bencoding.List{Items: items}, nil
}

func (p *Parser) parseString() (bencoding.Value, error) {
	if _, err := p.expect(token.StringLength); err != nil {
		return nil, err
	}

	if _, err := p.expect(token.StringSeperator); err != nil {
		return nil, err
	}

	valueToken, err := p.expect(token.String)
	if err != nil {
		return nil, err
	}

	return bencoding.String{Value: valueToken.Value}, nil
}

func (p *Parser) parseInteger() (bencoding.Value, error) {
	if _, err := p.expect(token.IntegerStart); err != nil {
		return nil, err
	}

	tok, err := p.expect(token.Integer)
	if err != nil {
		return nil, err
	}

	if _, err := p.expect(token.IntegerEnd); err != nil {
		return nil, err
	}

	value, err := strconv.ParseInt(tok.Value, 10, 64)
	if err != nil {
		return nil, err
	}

	return bencoding.Integer{Value: value}, nil
}
