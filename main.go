package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/eve-an/torrentd/pkg/hash"
	"github.com/eve-an/torrentd/pkg/parser"
	"github.com/eve-an/torrentd/pkg/torrent"
)

type args struct {
	file        string
	torrentFile string
}

func filePathExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !errors.Is(err, os.ErrNotExist)
}

func parseArgs() (args, error) {
	if len(os.Args) != 3 {
		return args{}, errors.New("unexpected number of arguments")
	}

	file, torrentFile := os.Args[1], os.Args[2]

	if !filePathExists(file) {
		return args{}, fmt.Errorf("file '%s' does not exist", file)
	}

	if !filePathExists(torrentFile) {
		return args{}, fmt.Errorf("torrentFile '%s' does not exist", file)
	}

	return args{file, torrentFile}, nil
}

func fail(err error) {
	fmt.Fprintln(os.Stderr, "error:", err)
	os.Exit(1)
}

func main() {
	args, err := parseArgs()
	if err != nil {
		fail(err)
	}

	file, err := os.Open(args.file)
	if err != nil {
		fail(err)
	}
	defer file.Close()

	torrentFile, err := os.Open(args.torrentFile)
	if err != nil {
		fail(err)
	}
	defer torrentFile.Close()

	infoDict, err := parser.ParseFromReader(bufio.NewReader(torrentFile))
	if err != nil {
		fail(err)
	}

	meta, err := torrent.NewTorrentMeta(infoDict)
	if err != nil {
		fail(err)
	}

	fileHashes, err := hash.GenerateFileHashes(bufio.NewReader(file), meta.PieceLength)
	if err != nil {
		fail(err)
	}

	verifier, err := torrent.NewTorrentVerifier(meta, fileHashes)
	if err != nil {
		fail(err)
	}

	fmt.Printf("Progress %f\n", verifier.Progress())
}
