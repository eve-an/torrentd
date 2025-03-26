package torrent

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"slices"

	"github.com/eve-an/torrentd/pkg/bencoding"
)

type TorrentMeta struct {
	PieceLength int
	Pieces      [][]byte
	Length      int64
	Name        string
}

func NewTorrentMeta(data bencoding.Value) (TorrentMeta, error) {
	dataDict, ok := data.(bencoding.Dict)
	if !ok {
		return TorrentMeta{}, errors.New("expected bencoding dict")
	}

	file, err := newMetaParser(dataDict)
	if err != nil {
		return TorrentMeta{}, err
	}

	var errs []error

	pieces, err := file.pieces()
	errs = append(errs, err)

	length, err := file.length()
	errs = append(errs, err)

	piecesLength, err := file.pieceLength()
	errs = append(errs, err)

	name, err := file.name()
	errs = append(errs, err)

	if err = errors.Join(errs...); err != nil {
		return TorrentMeta{}, err
	}

	return TorrentMeta{
		PieceLength: piecesLength,
		Length:      length,
		Name:        name,
		Pieces:      slices.Collect(slices.Chunk(pieces, sha1.Size)),
	}, nil
}

func (tm *TorrentMeta) String() string {
	return fmt.Sprintf("Name: %s, Length: %d, PieceLength: %d, Pieces: %d",
		tm.Name, tm.Length, tm.PieceLength, len(tm.Pieces))
}

type metaParser struct {
	data bencoding.Dict
}

func newMetaParser(data bencoding.Dict) (*metaParser, error) {
	info, found := data.Entities["info"]

	if !found {
		return nil, errors.New("no 'info' node found in bencoded data")
	}

	infoDict, ok := info.(bencoding.Dict)
	if !ok {
		return nil, errors.New("expected a dictionary for info node, got unexpected")
	}

	return &metaParser{
		data: infoDict,
	}, nil
}

func get[T bencoding.Value](t *metaParser, key string) (T, error) {
	value, found := t.data.Entities[key]
	if !found {
		return *new(T), fmt.Errorf("no '%s' key found in dict", key)
	}

	assertedValue, ok := value.(T)
	if !ok {
		return *new(T), fmt.Errorf("failure in type assertion for key '%s'", key)
	}

	return assertedValue, nil
}

func (t *metaParser) pieceLength() (int, error) {
	value, err := get[bencoding.Integer](t, "piece length")
	return int(value.Value), err
}

func (t *metaParser) pieces() ([]byte, error) {
	value, err := get[bencoding.String](t, "pieces")
	return []byte(value.Value), err
}

func (t *metaParser) name() (string, error) {
	value, err := get[bencoding.String](t, "name")
	return value.Value, err
}

func (t *metaParser) length() (int64, error) {
	value, err := get[bencoding.Integer](t, "length")
	return value.Value, err
}
