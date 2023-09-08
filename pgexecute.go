package main

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	_ "github.com/lib/pq"
)

func pgexecute(connection string, in string, out string) error {
	//	open postgresql connection
	db, err := sql.Open("postgres", connection)
	if err != nil {
		return err
	}
	defer db.Close()

	//	create output path
	err = os.MkdirAll(out, os.ModePerm)
	if err != nil {
		return err
	}

	//	get files from input path
	files, err := os.ReadDir(in)
	if err != nil {
		return err
	}
	for _, file := range files {
		if !file.Type().IsDir() && filepath.Ext(file.Name()) == ".sql" {
			log.Printf("file: %v", file.Name())

			//	read sql script
			buf, err := os.ReadFile(in + file.Name())
			if err != nil {
				return err
			}

			//	execute script
			_, err = db.Exec(string(buf))
			if err != nil {
				return err
			}
			log.Printf("executed")

			//	move file
			err = os.Rename(in+file.Name(), out+file.Name())
			if err != nil {
				return err
			}
		}
	}
	return nil
}
