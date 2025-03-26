package bencoding

import (
	"fmt"
	"maps"
	"slices"
	"strings"
)

// Here we use a tagged union because we dont have algebrais sum types :(
type Value interface {
	fmt.Stringer
	isValue()
}

type Integer struct {
	Value int64
}

type String struct {
	Value string
}

type List struct {
	Items []Value
}

type Dict struct {
	Entities map[string]Value
}

func (Integer) isValue() {}

func (i Integer) String() string {
	return fmt.Sprint(i.Value)
}

func (String) isValue() {}

func (s String) String() string {
	return `"` + s.Value + `"`
}

func (List) isValue() {}

func (l List) String() string {
	var sb strings.Builder

	sb.WriteString("[")
	for i, v := range l.Items {
		sb.WriteString(v.String())

		if i != len(l.Items)-1 {
			sb.WriteRune(',')
		}
	}
	sb.WriteString("]")

	return sb.String()
}

func (Dict) isValue() {}

func (d Dict) writeTo(sb *strings.Builder) {
	sb.WriteRune('{')

	keys := slices.Sorted(maps.Keys(d.Entities))
	for i, k := range keys {
		fmt.Fprintf(sb, `"%s":`, k)

		switch v := d.Entities[k].(type) {
		case Dict:
			v.writeTo(sb)
		default:
			sb.WriteString(v.String())
		}

		if i < len(keys)-1 {
			sb.WriteRune(',')
		}
	}

	sb.WriteRune('}')
}

func (d Dict) String() string {
	var sb strings.Builder
	d.writeTo(&sb)
	return sb.String()
}
