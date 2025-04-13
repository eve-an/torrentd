package status

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func openPanic(filePath string) *os.File {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	return file
}

func TestStatusService_CheckStatus(t *testing.T) {
	tests := []struct {
		name        string
		file        io.ReadCloser
		torrentFile io.ReadCloser
		want        Status
		wantErr     assert.ErrorAssertionFunc
	}{
		{
			"test sample torrent",
			openPanic("./fixture/sample.txt.part"),
			openPanic("./fixture/sample.torrent"),
			Status{
				Progress:        0.6666666666666666,
				TotalBlocks:     3,
				MissingBlocks:   1,
				CompletedBlocks: 2,
			},
			assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer tt.file.Close()
			defer tt.torrentFile.Close()

			s := NewStatusService()
			got, gotErr := s.CheckStatus(tt.file, tt.torrentFile)

			tt.wantErr(t, gotErr)
			assert.Equal(t, tt.want, got)
		})
	}
}
