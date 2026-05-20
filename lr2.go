package bmsdb

import (
	"os"
	"path/filepath"
	"strings"
)

type LR2Reader struct{}

func NewLR2Reader() *LR2Reader {
	return &LR2Reader{}
}

func (reader *LR2Reader) Score(query *QueryContext) ([]LR2Score, error) {
	if err := validateQuery(query); err != nil {
		return nil, err
	}
	db, err := openDatabase(query.path)
	if err != nil {
		return nil, err
	}

	var data []LR2Score
	if err := db.Find(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (reader *LR2Reader) Song(query *QueryContext) ([]LR2Song, error) {
	if err := validateQuery(query); err != nil {
		return nil, err
	}
	db, err := openDatabase(query.path)
	if err != nil {
		return nil, err
	}

	var data []LR2Song
	if err := db.Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

type LR2Scanner struct{}

func NewLR2Scanner() *LR2Scanner {
	return &LR2Scanner{}
}

func (scanner *LR2Scanner) ScanDirectory(dir string) (*LR2ScanResult, error) {
	if err := ValidateLR2Installation(dir); err != nil {
		return nil, err
	}

	database := filepath.Join(dir, "LR2files", "Database")

	ret := &LR2ScanResult{}
	song := filepath.Join(database, "song.db")
	if err := isExistedFile(song); err == nil {
		ret.Song = song
	}

	scoreFiles, err := scanLR2ScoreFolder(filepath.Join(database, "Score"))
	if err != nil {
		return nil, err
	}
	ret.Scores = scoreFiles
	return ret, nil
}

func ValidateLR2Installation(dir string) error {
	if err := isExistedDir(dir); err != nil {
		return err
	}

	lr2files := filepath.Join(dir, "LR2files")
	if err := isExistedDir(lr2files); err != nil {
		return err
	}

	database := filepath.Join(lr2files, "Database")

	if err := isExistedDir(database); err != nil {
		return err
	}

	return nil
}

func scanLR2ScoreFolder(dir string) ([]LR2PlayerScoreFile, error) {
	possibleFiles, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	ret := make([]LR2PlayerScoreFile, 0)
	for _, file := range possibleFiles {
		name := file.Name()
		if !strings.HasSuffix(name, ".db") {
			continue
		}
		ret = append(ret, LR2PlayerScoreFile{
			Name: name[:len(name)-3],
			Path: filepath.Join(dir, name),
		})
	}

	return ret, nil
}
