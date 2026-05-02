package bmsdb_test

import (
	"log"
	"testing"

	"github.com/Catizard/bmsdb"
)

func TestLR2FullOnRealFiles(t *testing.T) {
	dir := ".database/LR2"
	if err := isExistedDir(dir); err != nil {
		t.Skipf("%s is not an existed directory: %s", dir, err)
	}

	scanner := bmsdb.NewLR2Scanner()
	files, err := scanner.ScanDirectory(dir)
	if err != nil {
		t.Error(err)
	}

	reader := bmsdb.NewLR2Reader()

	songs, err := reader.Song(bmsdb.NewQueryContext(files.Song))
	if err != nil {
		t.Error(err)
	}

	log.Printf("song count: %d", len(songs))

	log.Printf("score files count: %d", len(files.Scores))
	for _, file := range files.Scores {
		log.Printf("score file: %s", file.Name)

		scores, err := reader.Score(bmsdb.NewQueryContext(file.Path))
		if err != nil {
			t.Error(err)
		}
		log.Printf("score count: %d", len(scores))
	}
}
