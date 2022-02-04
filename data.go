package main

import (
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
)

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

type NewJob struct {
	Position     string `form:"position"`
	Organization string `form:"organization"`
	Url          string `form:"url"`
	Description  string `form:"description"`
	Email        string `form:"email"`
}

func (newJob *NewJob) validate() ([]string, bool) {
	errs := []string{}
	if newJob.Position == "" {
		errs = append(errs, "Must provide a Position")
	}
	if newJob.Organization == "" {
		errs = append(errs, "Must provide a Organization")
	}
	if newJob.Url == "" && newJob.Description == "" {
		errs = append(errs, "Must provide either a Url or a Description")
	}
	// TODO: validate Url
	if newJob.Email == "" {
		errs = append(errs, "Must provide a Email")
	}
	// TODO: validate email
	return errs, len(errs) == 0
}

func (newJob *NewJob) saveToDB(db *sqlx.DB) (sql.Result, error) {
	query := "INSERT INTO jobs (position, organization, url, description, email) VALUES ($1, $2, $3, $4, $5)"

	params := []interface{}{
		newJob.Position,
		newJob.Organization,
		sql.NullString{
			String: newJob.Url,
			Valid:  newJob.Url != "",
		},
		sql.NullString{
			String: newJob.Description,
			Valid:  newJob.Description != "",
		},
		newJob.Email,
	}

	return db.Exec(query, params...)
}
