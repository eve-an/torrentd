package hash

import (
	"crypto/sha1"
	"errors"
	"io"
)

func readChunks(r io.Reader, chunkSize int) ([][]byte, error) {
	var chunks [][]byte

	buf := make([]byte, chunkSize)
	for {
		n, err := r.Read(buf)
		if n > 0 {
			chunk := make([]byte, n)
			copy(chunk, buf[:n])
			chunks = append(chunks, chunk)
		}

		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return nil, err
		}
	}

	return chunks, nil
}

func GenerateFileHashes(file io.Reader, chunkSize int) ([][]byte, error) {
	if chunkSize < 1 {
		return nil, errors.New("chunk size must be greater than 0")
	}

	fileChunks, err := readChunks(file, chunkSize)
	if err != nil {
		return nil, err
	}

	for i, chunk := range fileChunks {
		checkSum := sha1.Sum(chunk)
		fileChunks[i] = checkSum[:] // reuse fileChunks to save allocs
	}

	return fileChunks, nil
}
