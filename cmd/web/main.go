package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "http service address")
	dsn := flag.String("dsn", "web:goof@/snippetbox?parseTime=true","MYSQL data source name")
	flag.Parse()

	infoLog := log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal()
	}

	defer db.Close()

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
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dsn string)(*sql.DB, error){
	db, err := sql.Open("mysql",dsn)
	if err!= nil{
		return nil, err
	}
	if err = db.Ping(); err!=nil{
		return nil, err
	}
	return db, nil
}
