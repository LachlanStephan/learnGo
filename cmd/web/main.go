package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/rs/cors"
)

type config struct {
	addr      string
	staticDir string
}

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

var (
	cnf      config
	infoLog  = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Llongfile)
)

func setFlags() {
	flag.StringVar(&cnf.addr, "addr", ":4000", "HTTP network address")
	flag.StringVar(&cnf.staticDir, "static-dir", "./ui/static", "path to static assets")
	flag.Parse()
}

func main() {
	setFlags()

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	srv := &http.Server{
		Addr:     cnf.addr,
		ErrorLog: errorLog,
		Handler:  cors.Default().Handler(app.routes()),
	}

	infoLog.Printf("starting server on %s", cnf.addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
