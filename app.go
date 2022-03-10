package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal(err)
	}

	// test the connection to verify it actually works
	if err = db.Ping(); err != nil {
		log.Println(err)
	}

	go func() {
		ticker := time.NewTicker(time.Minute * 30)
		for {
			select {
			case <-ticker.C:
				if err = db.Ping(); err != nil {
					log.Println(err)
				} else {
					log.Println("Pinged DB successfully")
				}
			}
		}
	}()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err = db.Ping(); err != nil {
			log.Println(err)
		} else {
			log.Println("Pinged DB successfully")
		}
		w.Write([]byte("Pinging database..."))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
