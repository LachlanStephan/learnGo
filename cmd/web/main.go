package main

import (
	"database/sql"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/LachlanStephan/ls_server/internal/models"
	_ "github.com/go-sql-driver/mysql"

	"github.com/alexedwards/scs/mysqlstore"
	"github.com/alexedwards/scs/v2"
)

type config struct {
	addr      string
	staticDir string
	dsn       string
}

type application struct {
	errorLog       *log.Logger
	infoLog        *log.Logger
	users          *models.UserModel
	blogs          *models.BlogModel
	templateCache  map[string]*template.Template
	SessionManager *scs.SessionManager
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

func setSessionManager(db *sql.DB) *scs.SessionManager {
	sessionManager := scs.New()
	sessionManager.Store = mysqlstore.New(db)
	sessionManager.Lifetime = 12 * time.Hour
	return sessionManager
}

func main() {
	setFlags()

	db, err := openDB(cnf.dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	defer db.Close()

	templateCache, err := newTemplateCache()
	if err != nil {
		errorLog.Fatal(err)
	}

	sessionManager := setSessionManager(db)

	app := &application{
		errorLog:       errorLog,
		infoLog:        infoLog,
		users:          &models.UserModel{DB: db},
		blogs:          &models.BlogModel{DB: db},
		templateCache:  templateCache,
		SessionManager: sessionManager,
	}

	srv := &http.Server{
		Addr:           cnf.addr,
		MaxHeaderBytes: 524288,
		ErrorLog:       errorLog,
		Handler:        app.routes(),
		IdleTimeout:    time.Minute,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	infoLog.Printf("starting server on %s", cnf.addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}
