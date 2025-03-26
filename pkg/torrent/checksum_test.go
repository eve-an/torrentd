package torrent

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTorrentVerrifier_MissingBlocks(t *testing.T) {
	tests := []struct {
		name          string
		fileHashes    [][]byte
		torrentPieces [][]byte
		want          int
		wantErr       assert.ErrorAssertionFunc
	}{
		{
			name: "no missing blocks",
			fileHashes: [][]byte{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
			},
			torrentPieces: [][]byte{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
			},
			want:    0,
			wantErr: assert.NoError,
		},
		{
			name: "different values same chunks amount",
			fileHashes: [][]byte{
				{255, 2, 3, 4},
				{255, 6, 7, 8},
				{255, 10, 11, 12},
			},
			torrentPieces: [][]byte{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
			},
			want:    3,
			wantErr: assert.NoError,
		},
		{
			name: "different amount of hashes",
			fileHashes: [][]byte{
				{255, 6, 7, 8},
				{255, 10, 11, 12},
			},
			torrentPieces: [][]byte{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
			},
			want:    3,
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			to, err := NewTorrentVerifier(TorrentMeta{Pieces: tt.torrentPieces}, tt.fileHashes)

			if !tt.wantErr(t, err) {
				got := to.MissingBlocks()
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
