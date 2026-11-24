package main

import (
	"database/sql"
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

//go:embed web
var webFS embed.FS

func env(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}

func main() {
	dataDir := env("DATA_DIR", "data")
	addr := env("ADDR", "127.0.0.1:8080") // por defecto solo local
	if err := os.MkdirAll(dataDir, 0o755); err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("sqlite", filepath.Join(dataDir, "floreria.db")+"?_pragma=journal_mode(WAL)&_pragma=busy_timeout(5000)")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.SetMaxOpenConns(1) // SQLite: un escritor, simple y seguro

	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS ping (id INTEGER PRIMARY KEY, ts TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP)`); err != nil {
		log.Fatal(err)
	}

	sub, _ := fs.Sub(webFS, "web")
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.FS(sub)))
	mux.HandleFunc("/api/ping", func(w http.ResponseWriter, r *http.Request) {
		if _, err := db.Exec(`INSERT INTO ping DEFAULT VALUES`); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		var n int
		var ver string
		db.QueryRow(`SELECT COUNT(*) FROM ping`).Scan(&n)
		db.QueryRow(`SELECT sqlite_version()`).Scan(&ver)
		fmt.Fprintf(w, "Base de datos OK · SQLite %s · %d visitas guardadas", ver, n)
	})

	log.Printf("Escuchando en http://%s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
