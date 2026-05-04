package bmsdb_test

import (
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/Catizard/bmsdb"
	"github.com/glebarez/sqlite"
	"github.com/google/go-cmp/cmp"
	"github.com/rotisserie/eris"
	"gorm.io/gorm"
)

const MEMORY_DSN = "file::memory:?cache=shared"

func TestBeatorajaScanerSimple(t *testing.T) {
	tempDir := t.TempDir()

	songdata := filepath.Join(tempDir, "songdata.db")
	if _, err := os.Create(songdata); err != nil {
		t.Error(err)
	}
	songinfo := filepath.Join(tempDir, "songinfo.db")
	if _, err := os.Create(songinfo); err != nil {
		t.Error(err)
	}

	if err := os.Mkdir(filepath.Join(tempDir, "player"), 0777); err != nil {
		t.Error(err)
	}

	player1 := filepath.Join(tempDir, "player", "player1")

	if err := os.Mkdir(player1, 0777); err != nil {
		t.Error(err)
	}

	score := filepath.Join(player1, "score.db")
	if _, err := os.Create(score); err != nil {
		t.Error(err)
	}

	scorelog := filepath.Join(player1, "scorelog.db")
	if _, err := os.Create(scorelog); err != nil {
		t.Error(err)
	}

	scoredatalog := filepath.Join(player1, "scoredatalog.db")
	if _, err := os.Create(scoredatalog); err != nil {
		t.Error(err)
	}

	scanner := bmsdb.NewBeatorajaScanner()
	ret, err := scanner.ScanDirectory(tempDir)
	if err != nil {
		t.Error(err)
	}

	if ret.SongData != songdata {
		t.Errorf("expected %s, got %s", songdata, ret.SongData)
	}
	if ret.SongInfo != songinfo {
		t.Errorf("expected %s, got %s", songinfo, ret.SongInfo)
	}

	if len(ret.Players) != 1 {
		t.Errorf("players count expected %d, got %d", 1, len(ret.Players))
	}

	p := ret.Players[0]
	if p.Score != score {
		t.Errorf("expected %s, got %s", score, p.Score)
	}
	if p.ScoreLog != scorelog {
		t.Errorf("expected %s, got %s", scorelog, p.ScoreLog)
	}
	if p.ScoreDataLog != scoredatalog {
		t.Errorf("expected %s, got %s", scoredatalog, p.ScoreDataLog)
	}
}

func TestBeatorajaFullOnRealFiles(t *testing.T) {
	dir := ".database/beatoraja"
	if err := isExistedDir(dir); err != nil {
		t.Skipf("%s is not an existed directory: %s", dir, err)
	}

	scanner := bmsdb.NewBeatorajaScanner()
	files, err := scanner.ScanDirectory(dir)
	if err != nil {
		t.Error(err)
	}

	reader := bmsdb.NewBeatorajaReader()

	songDatas, err := reader.SongData(bmsdb.NewQueryContext(files.SongData))
	if err != nil {
		t.Error(err)
	}

	log.Printf("songdata count: %d", len(songDatas))

	songInfos, err := reader.SongInfo(bmsdb.NewQueryContext(files.SongInfo))
	if err != nil {
		t.Error(err)
	}

	log.Printf("songinfo count: %d", len(songInfos))

	// Players

	log.Printf("players count: %d", len(files.Players))
	for _, player := range files.Players {
		log.Printf("player: %s", player.Name)
		scores, err := reader.Score(bmsdb.NewQueryContext(player.Score))
		if err != nil {
			t.Error(err)
		}
		log.Printf("score count: %d", len(scores))

		scorelog, err := reader.ScoreLog(bmsdb.NewQueryContext(player.ScoreLog))
		if err != nil {
			t.Error(err)
		}
		log.Printf("scorelog count: %d", len(scorelog))

		scoredatalog, err := reader.ScoreDataLog(bmsdb.NewQueryContext(player.ScoreDataLog))
		if err != nil {
			t.Error(err)
		}
		log.Printf("scoredatalog count: %d", len(scoredatalog))
	}
}

func TestReadScoreSimple(t *testing.T) {
	score := bmsdb.BeatorajaScoreData{
		Sha256:     "8c8d47378d8afcc1e6595e72d30d73b6c034163910a783fd6b103b4509e763a1",
		Mode:       0,
		Clear:      6,
		Epg:        516,
		Lpg:        474,
		Egr:        316,
		Lgr:        303,
		Egd:        26,
		Lgd:        44,
		Ebd:        4,
		Lbd:        0,
		Epr:        0,
		Lpr:        6,
		Ems:        17,
		Lms:        12,
		Notes:      1689,
		Combo:      497,
		Minbp:      39,
		AvgJudge:   18333,
		PlayCount:  3,
		ClearCount: 3,
		Trophy:     "rnmh",
		Ghost:      "H4sIAAAAAAAAAGVVCxbCMAiDuPuf2a2QBDZ1PqQFQvgYkZn13O9-ZbSc-VKUfO5T7It5dDrsT_bJCYC-SPtzGgpcupghHQdHCepDFtl2DPVcSgMncsbbWdBbCcrgd7vA0SaekEBjuwi5v1C2jbIhOBEm3tBjHdFsvTqFxe1IQEBvf6SUAehapRTPTE28E6xiqoJLNXExPdP6uTnszfvA__DVUUkTUs5A-JAFuX4z5GSN223mayIfH5qXElVOxF3sU1M8SH_Ahft5VBdTmq29gGUK8FZOripNQfT5EFRFFlueMNoMLf2AVzlWVmuKLxfD4KDf-IJPdTcwGn16Ua_lbKHFjWs1vb8E1zO8W-IEXjUVA-qkry91t-bi3c1XuKUnTHWMx2T0Ph_hhdPz91pwboC5SjkGO1V5qkGZZ9PRrfrVAZ7FVPsOiZ6bQ9ixhrnNGbDgwTB7HTnSQlVDeYYVXCtVmpg96V5p8BledJPZEsB_g1nQ8oOzc_8yx29DmQYAAA==",
		Option:     0,
		Seed:       5699161,
		Random:     0,
		TimeStamp:  1748430415,
		State:      0,
		ScoreHash:  "0353657be2b8b81c09b1d50609628409075b02148d33799b54ff06900fffea3b596",
	}

	db, cleanup, err := createEmptyScoreDB()
	if err != nil {
		t.Error(err)
	}
	defer cleanup()

	if err := db.Create(score).Error; err != nil {
		t.Error(err)
	}

	reader := bmsdb.NewBeatorajaReader()
	scores, err := reader.Score(bmsdb.NewQueryContext(MEMORY_DSN))
	if err != nil {
		t.Error(err)
	}

	if len(scores) != 1 {
		t.Errorf("scores count expected to be 1, got %d", len(scores))
	}

	realScore := scores[0]

	if !cmp.Equal(score, realScore) {
		t.Error(cmp.Diff(score, realScore))
	}
}

func TestReadScoreLogSimple(t *testing.T) {
	scoreLog := bmsdb.BeatorajaScoreLog{
		Sha256:    "ce1fbc0148c4dda8df71e9eaa73285bf58666d11078f1decf1b40654586cb5cfe7d74b59bc583fce48ee4acef2aa8b7cc8f27b6038e9dcc31827ad6262b07a8ed0fa6cdf0fdadf0ab5b767a9e24c529ef7031dbeec216c0eeee07da2029850b543e3300d0d300b0863e6a7698d348b959f34308f6a038a1eaa3a659fdfa73bc6",
		Mode:      10010,
		Clear:     1,
		OldClear:  0,
		Score:     114,
		OldScore:  0,
		Combo:     93,
		OldCombo:  0,
		Minbp:     3966,
		OldMinbp:  2147483647,
		TimeStamp: 1713092975,
	}

	db, cleanup, err := createEmptyScoreLogDB()
	if err != nil {
		t.Error(err)
	}

	defer cleanup()

	if err := db.Create(scoreLog).Error; err != nil {
		t.Error(err)
	}

	reader := bmsdb.NewBeatorajaReader()
	logs, err := reader.ScoreLog(bmsdb.NewQueryContext(MEMORY_DSN))
	if err != nil {
		t.Error(err)
	}

	if len(logs) != 1 {
		t.Errorf("scorelog count expected 1, got %d", len(logs))
	}

	realLog := logs[0]
	if !cmp.Equal(scoreLog, realLog) {
		t.Error(cmp.Diff(scoreLog, realLog))
	}
}

func TestReadScoreDataLogSimple(t *testing.T) {
	scoreDataLog := bmsdb.BeatorajaScoreDataLog{
		Sha256:     "ce1fbc0148c4dda8df71e9eaa73285bf58666d11078f1decf1b40654586cb5cf",
		Mode:       0,
		Clear:      0,
		Epg:        20,
		Lpg:        6,
		Egr:        59,
		Lgr:        3,
		Egd:        13,
		Lgd:        0,
		Ebd:        1,
		Lbd:        0,
		Epr:        0,
		Lpr:        1,
		Ems:        1,
		Lms:        0,
		Notes:      1071,
		Combo:      93,
		Minbp:      971,
		AvgJudge:   908238,
		PlayCount:  3,
		ClearCount: 0,
		Trophy:     "",
		Ghost:      "H4sIAAAAAAAA_-3MWwrAQAhD0dzU_a95JkqhWyh4_BAfBBn0mBK3REuzcXe1uZm5vLLnI2ETkif3R-Zaa_3eAePfOqQvBAAA",
		Option:     1,
		Seed:       5319964,
		Random:     0,
		TimeStamp:  1713092971,
		State:      0,
		ScoreHash:  "035e55f63e7b2acaf95eddd865aa7445e32d0cae12c74508b10e731cc52dcb3eaed",
	}

	db, cleanup, err := createEmptyScoreDataLogDB()
	if err != nil {
		t.Error(err)
	}

	defer cleanup()

	if err := db.Create(scoreDataLog).Error; err != nil {
		t.Error(err)
	}

	reader := bmsdb.NewBeatorajaReader()
	logs, err := reader.ScoreDataLog(bmsdb.NewQueryContext(MEMORY_DSN))
	if err != nil {
		t.Error(err)
	}

	if len(logs) != 1 {
		t.Errorf("scoredatalog count expected 1, got %d", len(logs))
	}

	realLog := logs[0]
	if !cmp.Equal(scoreDataLog, realLog) {
		t.Error(cmp.Diff(scoreDataLog, realLog))
	}
}

func TestReadSongDataSimple(t *testing.T) {
	songData := bmsdb.BeatorajaSongData{
		Md5:        "82fbb000f431997c84cd25dd230471c9",
		Sha256:     "d9640554c61bf6040916c178f459bd1b6c70e95bc592fc047912449549babbe7",
		Title:      "Healing smile [HappyHardcore Style edit] (Fuckin' Endless 7)",
		SubTitle:   "本当にごめんなさい　ｄ(*´∀｀*)b",
		Genre:      "ENDLESS UNHAPPY HARDCORE",
		Artist:     "Yamajet(^^) vs Sphere(･∀･∀･) vs ctc(*>∀<)ﾉ←大体こいつのせい",
		SubArtist:  `<div align=\"center\"> Dosukoi(^^)　</div>`,
		Tag:        "",
		Path:       "/Volumes/Untitled/BMS Pack/BMS Complete Pack 24.1.1/Healing smile [HappyHardcore Style edit]/healing_smille_hhc_7exC.bme",
		Folder:     "cf92275f",
		StageFile:  "特になし！",
		Banner:     "これも特になし！</br>",
		BackBmp:    "",
		Preview:    "preview_music.ogg",
		Parent:     "e57519b8",
		Level:      0,
		Difficulty: 5,
		MaxBpm:     89,
		MinBpm:     89,
		Length:     42933370,
		Mode:       7,
		Judge:      100,
		Feature:    4,
		Content:    131,
		Date:       1389677922,
		Favorite:   0,
		AddDate:    1773488359,
		Notes:      2033111,
		ChartHash:  "1f8f4889fafe375ae8cee374cbe5b7f1ed094c031226fca8691e0d8ca90f0ec5",
	}

	db, cleanup, err := createEmptySongDataDB()
	if err != nil {
		t.Error(err)
	}
	defer cleanup()

	if err := db.Create(songData).Error; err != nil {
		t.Error(err)
	}

	reader := bmsdb.NewBeatorajaReader()
	songs, err := reader.SongData(bmsdb.NewQueryContext(MEMORY_DSN))
	if err != nil {
		t.Error(err)
	}

	if len(songs) != 1 {
		t.Errorf("songdata count expected 1, got %d", len(songs))
	}

	realSong := songs[0]
	if !cmp.Equal(songData, realSong) {
		t.Error(cmp.Diff(songData, realSong))
	}
}

func isExistedDir(dir string) error {
	if stat, err := os.Stat(dir); err != nil {
		if os.IsNotExist(err) {
			return eris.Errorf("no file exists at %s", dir)
		}
		return eris.Errorf("cannot stat file at %s", dir)
	} else if !stat.IsDir() {
		return eris.Errorf("directory path %s is a file", dir)
	}
	return nil
}

func createEmptyScoreDB() (*gorm.DB, func(), error) {
	db, err := gorm.Open(sqlite.Open(MEMORY_DSN))
	if err != nil {
		return nil, nil, err
	}

	if err := db.AutoMigrate(bmsdb.BeatorajaScoreData{}); err != nil {
		return nil, nil, err
	}

	return db, func() {
		if err := db.Migrator().DropTable(bmsdb.BeatorajaScoreData{}); err != nil {
			log.Printf("failed to clean up table: %s, tests could be wrong!", err)
		}
	}, nil
}

func createEmptyScoreLogDB() (*gorm.DB, func(), error) {
	db, err := gorm.Open(sqlite.Open(MEMORY_DSN))
	if err != nil {
		return nil, nil, err
	}

	if err := db.AutoMigrate(bmsdb.BeatorajaScoreLog{}); err != nil {
		return nil, nil, err
	}

	return db, func() {
		if err := db.Migrator().DropTable(bmsdb.BeatorajaScoreLog{}); err != nil {
			log.Printf("failed to clean up table: %s, tests could be wrong!", err)
		}
	}, nil
}

func createEmptyScoreDataLogDB() (*gorm.DB, func(), error) {
	db, err := gorm.Open(sqlite.Open(MEMORY_DSN))
	if err != nil {
		return nil, nil, err
	}

	if err := db.AutoMigrate(bmsdb.BeatorajaScoreDataLog{}); err != nil {
		return nil, nil, err
	}

	return db, func() {
		if err := db.Migrator().DropTable(bmsdb.BeatorajaScoreDataLog{}); err != nil {
			log.Printf("failed to clean up table: %s, tests could be wrong!", err)
		}
	}, nil
}

func createEmptySongDataDB() (*gorm.DB, func(), error) {
	db, err := gorm.Open(sqlite.Open(MEMORY_DSN))
	if err != nil {
		return nil, nil, err
	}

	if err := db.AutoMigrate(bmsdb.BeatorajaSongData{}); err != nil {
		return nil, nil, err
	}

	return db, func() {
		if err := db.Migrator().DropTable(bmsdb.BeatorajaSongData{}); err != nil {
			log.Printf("failed to clean up table: %s, tests could be wrong!", err)
		}
	}, nil
}
