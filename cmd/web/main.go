package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/LachlanStephan/ls_server/internal/models"
	_ "github.com/go-sql-driver/mysql"
)

type config struct {
	addr      string
	staticDir string
	dsn       string
}

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	users *models.UserModel
	blogs *models.BlogModel
}

var (
	cnf      config
	infoLog  = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Llongfile)
)

func setFlags() {
	flag.StringVar(&cnf.addr, "addr", ":8080", "HTTP network address")
	flag.StringVar(&cnf.staticDir, "static-dir", "./ui/static", "path to static assets")
	flag.StringVar(&cnf.dsn, "dsn", "root:@/ls_server?parseTime=true", "update this later - MySQL data source name")
	flag.Parse()
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	setFlags()

	db, err := openDB(cnf.dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	defer db.Close()

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		users:    &models.UserModel{DB: db},
		blogs:    &models.BlogModel{DB: db},
	}

	srv := &http.Server{
		Addr:     cnf.addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("starting server on %s", cnf.addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}
