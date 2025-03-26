package main

import (
	"os"

	"github.com/eve-an/torrentd/pkg/lexer"
	"github.com/eve-an/torrentd/pkg/parser"
)

func main() {
	file, err := os.ReadFile("./05720aa7b3005b23ac5f58e6240d0cdc0b605ce7.torrent")
	if err != nil {
		panic(err)
	}

	tokens, err := lexer.NewLexer().Lex(string(file))
	if err != nil {
		panic(err)
	}

	parser := parser.NewParser(tokens)

	_, _ = parser.Parse()
	if err != nil {
		panic(err)
	}

	// torrentFile := torrent.NewTorrentMeta(data.(bencoding.Dict))

	// fmt.Printf("Length: %d\n", torrentFile.Length())
	// fmt.Printf("Name: %s\n", torrentFile.Name())
	// fmt.Printf("Piece Length: %d\n", torrentFile.PieceLength())

	// otherFile, err := os.ReadFile("test.file")
	// if err != nil {
	// 	panic(err)
	// }
	//
	// shas := [][]byte{}
	// sha := sha1.New()
	// for chunk := range slices.Chunk(otherFile, int(torrentFile.PieceLength())) {
	// 	sha.Write(chunk)
	//
	// 	shas = append(shas, sha.Sum(nil))
	// }
	//
	// for chunk := range slices.Chunk(torrentFile.Pieces(), sha1.Size) {
	// }
}
