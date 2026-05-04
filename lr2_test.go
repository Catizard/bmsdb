package bmsdb_test

import (
	"log"
	"testing"

	"github.com/Catizard/bmsdb"
	"github.com/glebarez/sqlite"
	"github.com/google/go-cmp/cmp"
	"gorm.io/gorm"
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

func TestReadLR2SongSimple(t *testing.T) {
	lr2Song := bmsdb.LR2Song{
		MD5:        "6e4858fb2f3c29abf2f93a728a9c1ea2",
		Title:      "started",
		SubTitle:   " [A14]",
		Genre:      "J-Airy Pop",
		Artist:     "Ym1024 feat. lamie*",
		SubArtist:  "BGA : Mentalstock obj:ucc",
		Tag:        "",
		Path:       "BMS\\started\\start_a14_ucc.bme",
		Folder:     "ed78b2e4",
		Type:       0,
		StageFile:  "title.png",
		Banner:     "banner.png",
		BackBmp:    "",
		Parent:     "f59bfe65",
		Level:      10,
		Difficulty: 4,
		MaxBpm:     136,
		MinBpm:     136,
		Mode:       14,
		Judge:      2,
		LongNote:   0,
		Bga:        1,
		Random:     0,
		Date:       1306531794,
		Favorite:   0,
		Txt:        0,
		Karinotes:  1001,
		AddDate:    1537668786,
		ExLevel:    0,
	}

	db, cleanup, err := createEmptyLR2SongDB()
	if err != nil {
		t.Error(err)
	}

	defer cleanup()

	if err := db.Create(lr2Song).Error; err != nil {
		t.Error(err)
	}

	reader := bmsdb.NewLR2Reader()
	songs, err := reader.Song(bmsdb.NewQueryContext(MEMORY_DSN))
	if err != nil {
		t.Error(err)
	}

	if len(songs) != 1 {
		t.Errorf("LR2 song count expected 1, got %d", len(songs))
	}

	realSong := songs[0]
	if !cmp.Equal(lr2Song, realSong) {
		t.Error(cmp.Diff(lr2Song, realSong))
	}
}

func createEmptyLR2SongDB() (*gorm.DB, func(), error) {
	db, err := gorm.Open(sqlite.Open(MEMORY_DSN))
	if err != nil {
		return nil, nil, err
	}

	if err := db.AutoMigrate(bmsdb.LR2Song{}); err != nil {
		return nil, nil, err
	}

	return db, func() {
		if err := db.Migrator().DropTable(bmsdb.LR2Song{}); err != nil {
			log.Printf("failed to clean up table: %s, tests could be wrong!", err)
		}
	}, nil
}

func TestReadLR2ScoreSimple(t *testing.T) {
	lr2Score := bmsdb.LR2Score{
		MD5:        "8babf30e6f4fad6fd85fc94004bca86b",
		Clear:      1,
		Perfect:    589,
		Great:      359,
		Good:       116,
		Bad:        10,
		Poor:       8,
		TotalNotes: 1074,
		MaxCombo:   498,
		Minbp:      12,
		PlayCount:  1,
		ClearCount: 0,
		FailCount:  1,
		Rank:       6,
		Rate:       71,
		ClearDB:    0,
		OpHistory:  65795,
		ScoreHash:  "3f4f7224bd94451cb825667441b9e5df",
		Ghost:      "tScYSESEuSh2hEYYSh2GuTEstsq2THSEYShh3h4ETHXVh4GhESTEuh4GhGq5xXSESEtscESGh2cTMtsSGq3qXSEq2g3uqTETGsSEuq2XSEqr1tBEX@h@BghHhEtgh@YJhIq2gqYEqSF2cESEuq3usSGqg4TESGycsXTGXY2EsVIXh2ESEvXh2g3guSKgwUh2gBXBiGX@geY@qg2YEtYYYEwSGsSHSIstqq2@2ITcMvssSJsqVBGSKsSEXSJSEYXVEshP2EXhPXcP2PNvSJTHwXYYSETPEtwvwSh2Eq2r0q3SIuXYShgVPSPYPTEq6sq3h2ETF0wXBQP2EYEuttq6wyh2LuvSPPPPPPQED@cPPQg11tsSGXVhETMqSJXh4h2cY2EqXhgggYESghGcGcGSEZ",
		ClearSD:    0,
		ClearEX:    0,
		OpBest:     0,
		RSeed:      987,
		Complete:   1,
	}

	db, cleanup, err := createEmptyLR2UserDB()
	if err != nil {
		t.Fatal(err)
	}

	defer cleanup()

	if err := db.Create(&lr2Score).Error; err != nil {
		t.Error(err)
	}

	var scores []bmsdb.LR2Score
	if err := db.Find(&scores).Error; err != nil {
		t.Error(err)
	}

	if len(scores) != 1 {
		t.Errorf("expected 1 score, got %d", len(scores))
	}

	realScore := scores[0]

	realScore.RowID = 0
	lr2Score.RowID = 0

	if !cmp.Equal(lr2Score, realScore) {
		t.Error(cmp.Diff(lr2Score, realScore))
	}
}

func createEmptyLR2UserDB() (*gorm.DB, func(), error) {
	db, err := gorm.Open(sqlite.Open(MEMORY_DSN))
	if err != nil {
		return nil, nil, err
	}

	if err := db.AutoMigrate(bmsdb.LR2Score{}); err != nil {
		return nil, nil, err
	}

	return db, func() {
		if err := db.Migrator().DropTable(bmsdb.LR2Score{}); err != nil {
			log.Printf("failed to clean up table: %s, tests could be wrong!", err)
		}
	}, nil
}
