package main

import (
	"database/sql"
	"errors"
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

func getAllJobs(db *sqlx.DB) ([]Job, error) {
	var jobs []Job

	err := db.Select(&jobs, "SELECT * FROM jobs")
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return jobs, err
	}

	return jobs, nil
}

func getJob(id string, db *sqlx.DB) (Job, error) {
	var job Job

	err := db.Get(&job, "SELECT * FROM jobs WHERE id = $1", id)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return job, err
	}

	return job, nil
}

type NewJob struct {
	Position     string `form:"position"`
	Organization string `form:"organization"`
	Url          string `form:"url"`
	Description  string `form:"description"`
	Email        string `form:"email"`
}

func (newJob *NewJob) validate() map[string]string {
	errs := make(map[string]string)

	if newJob.Position == "" {
		errs["position"] = "Must provide a Position"
	}

	if newJob.Organization == "" {
		errs["organization"] = "Must provide a Organization"
	}

	if newJob.Url == "" && newJob.Description == "" {
		errs["url"] = "Must provide either a Url or a Description"
	}

	// TODO: validate Url

	if newJob.Email == "" {
		// TODO: validate email
		errs["email"] = "Must provide a valid Email"
	}

	return errs
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
