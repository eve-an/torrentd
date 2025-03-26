package torrent_test

import (
	"testing"

	"github.com/eve-an/torrentd/pkg/bencoding"
	"github.com/eve-an/torrentd/pkg/torrent"
	"github.com/stretchr/testify/assert"
)

func TestNewTorrentMeta(t *testing.T) {
	tests := []struct {
		name    string
		data    bencoding.Dict
		want    torrent.TorrentMeta
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "successfull creation with valid data dict",
			data: bencoding.Dict{Entities: map[string]bencoding.Value{
				"info": bencoding.Dict{Entities: map[string]bencoding.Value{
					"pieces":       bencoding.String{Value: "012345"},
					"piece length": bencoding.Integer{Value: 1},
					"name":         bencoding.String{Value: "test"},
					"length":       bencoding.Integer{Value: 2},
				}},
			}},
			want: torrent.TorrentMeta{
				PieceLength: 1,
				Length:      2,
				Pieces:      [][]byte{[]byte("012345")},
				Name:        "test",
			},
			wantErr: assert.NoError,
		},
		{
			name: "info key is missing!",
			data: bencoding.Dict{Entities: map[string]bencoding.Value{
				"not my info key :(": bencoding.Dict{Entities: map[string]bencoding.Value{
					"pieces":       bencoding.String{Value: "012345"},
					"piece length": bencoding.Integer{Value: 1},
					"name":         bencoding.String{Value: "test"},
					"length":       bencoding.Integer{Value: 2},
				}},
			}},
			want:    torrent.TorrentMeta{},
			wantErr: assert.Error,
		},
		{
			name: "pieces key is missing",
			data: bencoding.Dict{Entities: map[string]bencoding.Value{
				"info": bencoding.Dict{Entities: map[string]bencoding.Value{
					"piece length": bencoding.Integer{Value: 1},
					"name":         bencoding.String{Value: "test"},
					"length":       bencoding.Integer{Value: 2},
				}},
			}},
			want:    torrent.TorrentMeta{},
			wantErr: assert.Error,
		},
		{
			name: "piece length key is missing",
			data: bencoding.Dict{Entities: map[string]bencoding.Value{
				"info": bencoding.Dict{Entities: map[string]bencoding.Value{
					"pieces": bencoding.String{Value: "012345"},
					"name":   bencoding.String{Value: "test"},
					"length": bencoding.Integer{Value: 2},
				}},
			}},
			want:    torrent.TorrentMeta{},
			wantErr: assert.Error,
		},
		{
			name: "name key is missing",
			data: bencoding.Dict{Entities: map[string]bencoding.Value{
				"info": bencoding.Dict{Entities: map[string]bencoding.Value{
					"pieces":       bencoding.String{Value: "012345"},
					"piece length": bencoding.Integer{Value: 1},
					"length":       bencoding.Integer{Value: 2},
				}},
			}},
			want:    torrent.TorrentMeta{},
			wantErr: assert.Error,
		},
		{
			name: "length key is missing",
			data: bencoding.Dict{Entities: map[string]bencoding.Value{
				"info": bencoding.Dict{Entities: map[string]bencoding.Value{
					"pieces":       bencoding.String{Value: "012345"},
					"piece length": bencoding.Integer{Value: 1},
					"name":         bencoding.String{Value: "test"},
				}},
			}},
			want:    torrent.TorrentMeta{},
			wantErr: assert.Error,
		},
		{
			name: "wrong value types",
			data: bencoding.Dict{Entities: map[string]bencoding.Value{
				"info": bencoding.Dict{Entities: map[string]bencoding.Value{
					"pieces":       bencoding.Integer{Value: 123},
					"piece length": bencoding.String{Value: "1"},
					"name":         bencoding.Integer{Value: 0},
					"length":       bencoding.String{Value: "2"},
				}},
			}},
			want:    torrent.TorrentMeta{},
			wantErr: assert.Error,
		},
		{
			name: "info is not a dict",
			data: bencoding.Dict{Entities: map[string]bencoding.Value{
				"info": bencoding.String{Value: "test"},
			}},
			want:    torrent.TorrentMeta{},
			wantErr: assert.Error,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := torrent.NewTorrentMeta(tt.data)

			tt.wantErr(t, gotErr)
			assert.Equal(t, tt.want, got)
		})
	}
}
