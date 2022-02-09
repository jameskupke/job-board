package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"time"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	config, err := LoadConfig()
	if err != nil {
		panic(err)
	}

	gin.SetMode(config.Env)

	// migrate the db on startup
	m, err := migrate.New("file://sql", config.DatabaseURL)
	if err != nil {
		panic(err)
	}

	m.Up()

	db, err := sqlx.Open("postgres", config.DatabaseURL)
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			_, err := db.Exec("DELETE FROM jobs WHERE published_at < NOW() - INTERVAL '30 DAYS'")
			if err != nil {
				fmt.Printf("error clearing old jobs: %s\n", err.Error())
			}
			time.Sleep(1 * time.Hour)
		}
	}()

	ctrl := &Controller{DB: db, Config: config}

	router := gin.Default()
	router.SetTrustedProxies(nil)

	sessionOpts := sessions.Options{
		Path:     "/",
		MaxAge:   24 * 60, // 1 day
		Secure:   config.Env != "debug",
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}

	sessionStore := cookie.NewStore([]byte(config.AppSecret))
	sessionStore.Options(sessionOpts)
	router.Use(sessions.Sessions("mysession", sessionStore))

	router.Static("/assets", "assets")
	router.HTMLRender = renderer()

	router.GET("/", ctrl.Index)
	router.GET("/new", ctrl.NewJob)
	router.POST("/jobs", ctrl.CreateJob)
	router.GET("/jobs/:id", ctrl.ViewJob)

	authorized := router.Group("/")
	authorized.Use(requireAuth(db, config.AppSecret))
	{
		authorized.GET("/jobs/:id/edit", ctrl.EditJob)
		authorized.POST("/jobs/:id", ctrl.UpdateJob)
	}

	router.Run()
}

func renderer() multitemplate.Renderer {
	funcMap := template.FuncMap{
		"formatAsDate":          formatAsDate,
		"formatAsRfc3339String": formatAsRfc3339String,
	}

	r := multitemplate.NewRenderer()
	r.AddFromFilesFuncs("index", funcMap, "./templates/base.html", "./templates/index.html")
	r.AddFromFilesFuncs("new", funcMap, "./templates/base.html", "./templates/new.html")
	r.AddFromFilesFuncs("edit", funcMap, "./templates/base.html", "./templates/edit.html")
	r.AddFromFilesFuncs("view", funcMap, "./templates/base.html", "./templates/view.html")

	return r
}

func requireAuth(db *sqlx.DB, secret string) func(*gin.Context) {
	return func(ctx *gin.Context) {
		jobID := ctx.Param("id")
		job, err := getJob(jobID, db)
		if err != nil {
			panic(err) // TODO: handle!
		}

		token := ctx.Query("token")
		expected := signatureForJob(job, secret)

		fmt.Printf("token: %s\n", expected)

		if token != expected {
			ctx.AbortWithStatus(403)
			return
		}
	}
}
