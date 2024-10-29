package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "http service address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}




	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("listening on %v", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
