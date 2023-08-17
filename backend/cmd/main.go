package main

import (
	"errors"
	"flag"
	"log"
	"os"

	server "github.com/codescalersinternships/ToDoApp-Rodina/backend/internal"
	"github.com/gin-gonic/gin"
)

// ErrDBPathNotFound is returned if the user doesn't mention the path of the database
var ErrDBPathNotFound = errors.New("database path must be provided after the -db flag")

func main() {

	var dbPath string

	flag.StringVar(&dbPath, "db", "./todoapp.db", "mention the filepath of the database")

	flag.Parse()

	if dbPath == "" {
		log.Fatal(ErrDBPathNotFound)
	}

	db, err := server.ConnectToDB(dbPath)

	if err != nil {
		log.Fatal(err)
	}

	app := server.NewApp(db)

	// set gin mode from environment variable
	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "" {
		ginMode = gin.ReleaseMode
	}

	if err = app.Run(ginMode); err != nil {

		db.Close()
		log.Fatal(err)
	}

	db.Close()

}
