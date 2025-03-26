package torrent

import (
	"fmt"
	"slices"
)

type TorrentVerifier struct {
	fileHashes [][]byte
	torrent    TorrentMeta
}

func NewTorrentVerifier(meta TorrentMeta, fileHashes [][]byte) (*TorrentVerifier, error) {
	if len(meta.Pieces) != len(fileHashes) {
		return nil, fmt.Errorf("amount of hashes must be the same: torrent meta '%d' - file '%d'", len(meta.Pieces), len(fileHashes))
	}

	return &TorrentVerifier{
		torrent:    meta,
		fileHashes: fileHashes,
	}, nil
}

func (t *TorrentVerifier) IsCompleted() bool {
	return t.MissingBlocks() == 0
}

func (t *TorrentVerifier) Progress() float64 {
	total := len(t.torrent.Pieces)
	missing := t.MissingBlocks()

	return float64(total-missing) / float64(total)
}

func (t *TorrentVerifier) MissingBlocks() int {
	missing := 0

	for i, piece := range t.torrent.Pieces {
		if !slices.Equal(piece, t.fileHashes[i]) {
			missing++
		}
	}

	return missing
}
