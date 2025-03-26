package token

type Type = int

const (
	IntegerStart Type = iota
	Integer
	IntegerEnd
	StringLength
	StringSeperator
	String
	ListStart
	ListEnd
	DictionaryStart
	DictionaryEnd
)

type Token struct {
	Value string
	Type  Type
}
