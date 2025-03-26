package bencoding

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDict_String(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		d    Dict
		want string
	}{
		{
			name: "simple 1 level dict",
			d: Dict{Entities: map[string]Value{
				"a": Integer{Value: 1},
				"b": String{Value: "2"},
			}},
			want: `{"a":1,"b":"2"}`,
		},
		{
			name: "simple 2 level dict",
			d: Dict{Entities: map[string]Value{
				"a": Dict{Entities: map[string]Value{
					"b": Integer{Value: 1},
					"c": Integer{Value: 2},
				}},
			}},
			want: `{"a":{"b":1,"c":2}}`,
		},
		{
			name: "empty dict",
			d:    Dict{Entities: map[string]Value{}},
			want: `{}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.d.String()

			assert.Equal(t, tt.want, got)
		})
	}
}
