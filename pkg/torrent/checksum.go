package torrent

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"slices"
)

type TorrentVerifier struct {
	fileHashes [][]byte
	torrent    TorrentMeta
}

func generateFileHashes(file []byte, chunkSize int) ([][]byte, error) {
	if chunkSize < 1 {
		return nil, errors.New("chunk size must be greater than 0")
	}

	fileHashes := make([][]byte, 0, len(file)/chunkSize)
	for chunk := range slices.Chunk(file, chunkSize) {
		checkSum := sha1.Sum(chunk)
		fileHashes = append(fileHashes, checkSum[:])
	}

	return fileHashes, nil
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

// func NewTorrentVerifier(file, torrentFile io.Reader) (*TorrentVerifier, error) {
// 	torrentData, err := parser.ParseFromReader(torrentFile)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	meta, err := NewTorrentMeta(torrentData)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	fileData, err := io.ReadAll(file)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	fileHashes, err := generateFileHashes(fileData, meta.PieceLength)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	if len(fileHashes) != len(meta.Pieces) {
// 		return nil, fmt.Errorf("different amount of file hashes and torrent hashes: file hash '%d' - torrent hash '%d'", len(fileHashes), len(meta.Pieces))
// 	}
//
// 	return &TorrentVerifier{
// 		file:       fileData,
// 		torrent:    meta,
// 		fileHashes: fileHashes,
// 	}, nil
// }

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
