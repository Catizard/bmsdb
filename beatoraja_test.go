package bmsdb_test

import (
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/Catizard/bmsdb"
	"github.com/rotisserie/eris"
)

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
