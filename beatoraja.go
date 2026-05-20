package bmsdb

import (
	"os"
	"path/filepath"

	"github.com/rotisserie/eris"
)

type BeatorajaReader struct{}

func NewBeatorajaReader() *BeatorajaReader {
	return &BeatorajaReader{}
}

func (reader *BeatorajaReader) SongData(query *QueryContext) ([]BeatorajaSongData, error) {
	if err := validateQuery(query); err != nil {
		return nil, err
	}
	db, err := openDatabase(query.path)
	if err != nil {
		return nil, err
	}

	var data []BeatorajaSongData
	if err := db.Find(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (reader *BeatorajaReader) SongInfo(query *QueryContext) ([]BeatorajaSongInfo, error) {
	if err := validateQuery(query); err != nil {
		return nil, err
	}
	db, err := openDatabase(query.path)
	if err != nil {
		return nil, err
	}

	var data []BeatorajaSongInfo
	if err := db.Find(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

// Score returns the data in score.db
func (reader *BeatorajaReader) Score(query *QueryContext) ([]BeatorajaScoreData, error) {
	if err := validateQuery(query); err != nil {
		return nil, err
	}
	db, err := openDatabase(query.path)
	if err != nil {
		return nil, err
	}

	var data []BeatorajaScoreData
	if query.after != nil {
		db = db.Where("date > ?", *query.after)
	}
	if err := db.Find(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

// ScoreLog returns the data in scorelog.db
func (reader *BeatorajaReader) ScoreLog(query *QueryContext) ([]BeatorajaScoreLog, error) {
	if err := validateQuery(query); err != nil {
		return nil, err
	}

	db, err := openDatabase(query.path)
	if err != nil {
		return nil, err
	}

	var data []BeatorajaScoreLog
	if query.after != nil {
		db = db.Where("date > ?", *query.after)
	}
	if err := db.Find(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

// ScoreDataLog returns the data in scoredatalog.db
func (reader *BeatorajaReader) ScoreDataLog(query *QueryContext) ([]BeatorajaScoreDataLog, error) {
	if err := validateQuery(query); err != nil {
		return nil, err
	}

	db, err := openDatabase(query.path)
	if err != nil {
		return nil, err
	}

	var data []BeatorajaScoreDataLog
	if query.after != nil {
		db = db.Where("date > ?", *query.after)
	}
	if err := db.Find(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

type BeatorajaScanner struct{}

func NewBeatorajaScanner() *BeatorajaScanner {
	return &BeatorajaScanner{}
}

func (scanner *BeatorajaScanner) ScanDirectory(dir string) (*BeatorajaScanResult, error) {
	if err := ValidateBeatorajaInstallation(dir); err != nil {
		return nil, eris.Wrap(err, "not a valid beatoraja installation")
	}

	ret := &BeatorajaScanResult{}
	ret.SongData = filepath.Join(dir, "songdata.db")
	ret.SongInfo = filepath.Join(dir, "songinfo.db")

	player := filepath.Join(dir, "player")
	players, err := scanBeatorajaUserFolder(player)
	if err != nil {
		return nil, eris.Wrap(err, "scan player folder")
	}
	ret.Players = players
	return ret, nil
}

func ValidateBeatorajaInstallation(dir string) error {
	if err := isExistedDir(dir); err != nil {
		return err
	}

	if err := isExistedFile(filepath.Join(dir, "songdata.db")); err != nil {
		return err
	}

	if err := isExistedFile(filepath.Join(dir, "songinfo.db")); err != nil {
		return err
	}

	if err := isExistedDir(filepath.Join(dir, "player")); err != nil {
		return err
	}

	return nil
}

func scanBeatorajaUserFolder(dir string) ([]BeatorajaPlayerFolder, error) {
	possibleFolders, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	players := make([]BeatorajaPlayerFolder, 0)
	for _, entry := range possibleFolders {
		if !entry.IsDir() {
			continue
		}

		ret := BeatorajaPlayerFolder{}

		folder := filepath.Join(dir, entry.Name())
		children, err := os.ReadDir(folder)
		if err != nil {
			continue
		}

		for _, entry := range children {
			if entry.IsDir() {
				continue
			}
			switch entry.Name() {
			case "score.db":
				ret.Score = filepath.Join(folder, entry.Name())
			case "scorelog.db":
				ret.ScoreLog = filepath.Join(folder, entry.Name())
			case "scoredatalog.db":
				ret.ScoreDataLog = filepath.Join(folder, entry.Name())
			}
		}

		players = append(players, ret)
	}
	return players, nil
}
