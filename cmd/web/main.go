package main

import (
	"database/sql"
	"flag"
	"log"
	"time"

	"github.com/sum28it/snippetBox/pkg/models"

	"github.com/alexedwards/scs"
	_ "github.com/lib/pq"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", "postgres://snippet_box:snippet@localhost/snippet_box?sslmode=disable", "PostgreSQL DSN")
	htmlDir := flag.String("html-dir", "./ui/html", "Path to HTML templates")
	secret := flag.String("secret", "s6Ndh+pPbnzHbS*+9Pk8qGWhTzbpa@ge", "Secret key")
	staticDir := flag.String("static-dir", "./ui/static", "Path to static assets directory")
	tlsCert := flag.String("tls-cert", "./tls/cert.pem", "Path to TLS certificate")
	tlsKey := flag.String("tls-key", "./tls/key.pem", "Path to TLS key")

	flag.Parse()

	db := connect(*dsn)
	defer db.Close()

	sessionManager := scs.NewCookieManager(*secret)
	sessionManager.Lifetime(12 * time.Hour)
	sessionManager.Persist(true)

	app := &App{
		Addr:      *addr,
		Database:  &models.Database{db},
		HTMLDir:   *htmlDir,
		Sessions:  sessionManager,
		StaticDir: *staticDir,
		TLSCert:   *tlsCert,
		TLSKey:    *tlsKey,
	}

	app.RunServer()
}

func connect(dsn string) *sql.DB {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}
