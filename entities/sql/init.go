package sql

import (
	"database/sql"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	_ "github.com/lib/pq"
	"github.com/ribeirosaimon/aergia-utils/logs"
)

func CreateTableDatabase(conn string) error {
	return initTableDatabase(conn, nil)
}

// MockCreateTableDatabase I only create received table
func MockCreateTableDatabase(conn string, tables map[string]bool) error {
	return initTableDatabase(conn, tables)
}

func initTableDatabase(connStr string, tables map[string]bool) error {

	_, file, _, ok := runtime.Caller(0)
	if !ok {
		logs.ERROR.Message("not found files")
	}

	absPath := filepath.Dir(file)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		logs.ERROR.Message(err.Error())
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		logs.ERROR.Message(err.Error())
	}

	logs.LOG.Message("Connected to database")

	if err = filepath.Walk(absPath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".sql") {

			file, err := os.Open(path)
			if tables != nil {
				fileName := filepath.Base(path)
				if _, ok = tables[fileName]; !ok {
					return nil
				}
			}
			defer file.Close()
			logs.WARN.Message(fmt.Sprintf(".sql found: %s", path))

			if err != nil {
				logs.WARN.Message(fmt.Sprintf("error open file: %s", path))
				return err
			}

			sqlContent, err := io.ReadAll(file)
			if err != nil {
				logs.ERROR.Message(err.Error())
				return err
			}

			_, err = db.Exec(string(sqlContent))
			if err != nil {
				logs.ERROR.Message(err.Error())
				return err
			}
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}
