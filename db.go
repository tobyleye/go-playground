package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

type Album struct {
	Id     string
	Title  string
	Artist string
	Price  float32
}

func main() {
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "root",
		DBName:               "recordings",
		Net:                  "tcp",
		Addr:                 "localhost:3307",
		AllowNativePasswords: true,
	}

	formattedConfig := cfg.FormatDSN()
	fmt.Println(formattedConfig)
	var err error

	db, err := sql.Open("mysql", formattedConfig)

	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()

	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Db connected")

	albums, err := getAllAlbums(db)

	if err != nil {
		fmt.Println("error fetching albums")
	} else {
		fmt.Println("albums = ", albums)
	}

}

func getAllAlbums(db *sql.DB) ([]Album, error) {

	rows, err := db.Query("select * from album")
	albums := []Album{}

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	for rows.Next() {
		var album Album
		rowErr := rows.Scan(&album.Id, &album.Title, &album.Artist, &album.Price)
		if rowErr == nil {
			albums = append(albums, album)
		}
	}

	return albums, err

}
