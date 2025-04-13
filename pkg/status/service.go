package status

import (
	"io"

	"github.com/eve-an/torrentd/pkg/hash"
	"github.com/eve-an/torrentd/pkg/parser"
	"github.com/eve-an/torrentd/pkg/torrent"
)

type StatusService struct{}

func NewStatusService() *StatusService {
	return &StatusService{}
}

type Status struct {
	Progress        float64
	TotalBlocks     uint
	MissingBlocks   uint
	CompletedBlocks uint
}

func (s *StatusService) CheckStatus(file, torrentFile io.Reader) (Status, error) {
	infoDict, err := parser.ParseFromReader(torrentFile)
	if err != nil {
		return Status{}, err
	}

	meta, err := torrent.NewTorrentMeta(infoDict)
	if err != nil {
		return Status{}, err
	}

	fileHashes, err := hash.GenerateFileHashes(file, meta.PieceLength)
	if err != nil {
		return Status{}, err
	}

	verifier, err := torrent.NewTorrentVerifier(meta, fileHashes)
	if err != nil {
		return Status{}, err
	}

	missingBlocks := verifier.MissingBlocks()
	totalBlocks := verifier.TotalBlocks()

	return Status{
		Progress:        verifier.Progress(),
		MissingBlocks:   uint(missingBlocks),
		CompletedBlocks: uint(totalBlocks - missingBlocks),
		TotalBlocks:     uint(totalBlocks),
	}, nil
}
