package sql

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/lib/pq"
	"github.com/ribeirosaimon/aergia-utils/logs"
)

func initTableDatabase(connStr string) error {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to database")
	root, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("path:", root)

	if err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), ".sql") {
			logs.WARN.Message(fmt.Sprintf(".sql found: %s", path))

			file, err := os.Open(path)
			defer file.Close()

			if err != nil {
				logs.WARN.Message(fmt.Sprintf("Arquivo .sql encontrado: %s", path))
				return err
			}

			sqlContent, err := ioutil.ReadAll(file)
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
