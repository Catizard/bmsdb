package bmsdb

import (
	"os"

	"github.com/glebarez/sqlite"
	"github.com/rotisserie/eris"
	"gorm.io/gorm"
)

const READONLY_PARAMETER = "?open_mode=1"

func isExistedFile(fp string) error {
	if stat, err := os.Stat(fp); err != nil {
		if os.IsNotExist(err) {
			return eris.Errorf("no file exists at %s", fp)
		}
		return eris.Errorf("cannot stat file at %s", fp)
	} else if stat.IsDir() {
		return eris.Errorf("file path %s is a directory", fp)
	}
	return nil
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

func openDatabase(fp string) (*gorm.DB, error) {
	dsn := fp + READONLY_PARAMETER
	return gorm.Open(sqlite.Open(dsn))
}

func validateQuery(query *QueryContext) error {
	if query == nil {
		return eris.Errorf("query cannot be nil")
	}
	if query.path == "" {
		return eris.Errorf("query path cannot be empty")
	}

	if err := isExistedFile(query.path); err != nil {
		return eris.Wrap(err, "query path is not a valid file")
	}

	return nil
}
