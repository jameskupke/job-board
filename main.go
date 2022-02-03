package main

import (
	"fmt"
	"html/template"
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

	// TODO: better management of this seed data
	db.MustExec("INSERT INTO jobs (position, organization, url, email) VALUES ('pos1', 'org1', 'https://sethetter.com', 'seth@devict.org');")

	ctrl := &Controller{DB: db}

	router := gin.Default()

	router.SetFuncMap(template.FuncMap{
		"formatAsDate":          formatAsDate,
		"formatAsRfc3339String": formatAsRfc3339String,
	})
	router.LoadHTMLGlob("templates/*.html")
	router.Static("/assets", "assets")

	router.GET("/", ctrl.Home)
	router.GET("/new", ctrl.NewJob)

	router.Run()
}

// routes
// -------------------------------------

type Controller struct {
	DB *sqlx.DB
}

func (ctrl *Controller) Home(ctx *gin.Context) {
	var jobs []Job

	if err := ctrl.DB.Select(&jobs, "SELECT * FROM jobs"); err != nil {
		// TODO: handle error properly
		log.Fatal(err)
	}

	ctx.HTML(200, "index.html", gin.H{"jobs": jobs, "noJobs": len(jobs) == 0})
}

func (ctrl *Controller) NewJob(ctx *gin.Context) {
	ctx.HTML(200, "new.html", gin.H{})
}

// data
// -------------------------------------

type Job struct {
	ID           string         `db:"id"`
	Position     string         `db:"position"`
	Organization string         `db:"organization"`
	Url          sql.NullString `db:"url"`
	Description  sql.NullString `db:"description"`
	Email        string         `db:"email"`
	PublishedAt  time.Time      `db:"published_at"`
}

// template helpers
// ------------------------------------
func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d/%02d/%02d", year, month, day)
}

func formatAsRfc3339String(t time.Time) string {
	return t.Format(time.RFC3339)
}
