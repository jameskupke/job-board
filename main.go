package main

import (
	"log"
	"os"
	"time"

	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	// migrate the db on startup
	m, err := migrate.New("file://sql", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	m.Up()

	// connect to db
	db, err := sqlx.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	db.MustExec("INSERT INTO jobs (position, organization, url, email) VALUES ('pos1', 'org1', 'https://sethetter.com', 'seth@devict.org');")

	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")
	router.Static("/assets", "assets")

	router.GET("/", homeRoute)

	router.Run()
}

// routes
// -------------------------------------

func homeRoute(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{})
}

// data
// -------------------------------------

type Job struct {
	Position     string         `db:"position"`
	Organization string         `db:"organization"`
	Url          sql.NullString `db:"url"`
	Description  sql.NullString `db:"description"`
	Email        string         `db:"email"`
	PublishedAt  time.Time      `db:"published_at"`
}
