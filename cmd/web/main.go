package main

import (
	"database/sql"
	"flag"
	"glyst/internal/models"
	"html/template"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	logger *slog.Logger
	glysts *models.GlystModel
	templateCache map[string]*template.Template
}


func openDB(dsn string) (*sql.DB, error){
	db, err:= sql.Open("mysql", dsn)
	
	if err!=nil{
		return nil, err
	}

	err = db.Ping()

	if err !=nil{
		db.Close()
		return nil, err

	}
	return db, err
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", "web:pass@tcp(localhost:3306)/glyst?parseTime=true", "MySQL data source name")
	flag.Parse()

	// Use the slog.New() function to initialize a new structured logger, which
	// writes to the standard out stream and uses the default settings
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	}))

	
	db,err:=openDB(*dsn)
	if err!=nil{
		logger.Error(err.Error())
		os.Exit(1)	
	}


	defer db.Close()


	templateCache, err := newTemplateCache()

	if err!=nil{
	 	logger.Error(err.Error())	
		os.Exit(1)	
	}



	// Initialize a new instance of our application struct, containing the
	// dependencies (for now, just the structured logger).
	app := &application{
		logger: logger,
		glysts: &models.GlystModel{DB: db},
		templateCache: templateCache,	
	}

	// Print a log message to say that the server is starting.
	logger.Info("starting server", "addr", *addr)

	// Use the http.ListenAndServe() function to start a new web server. We pass in
	// two parameters: the TCP network address to listen on (in this case ":4000")
	// and the servemux we just created. If http.ListenAndServe() returns an error
	// we use the log.Fatal() function to log the error message and exit. Note
	// that any error returned by http.ListenAndServe() is always non-nil.
	err = http.ListenAndServe(*addr, app.routes())
	if err != nil {
		logger.Error("Server error", "err", err.Error())
		os.Exit(1)
	}
}
