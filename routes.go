package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type Controller struct {
	DB *sqlx.DB
}

func (ctrl *Controller) Index(ctx *gin.Context) {
	var jobs []Job

	if err := ctrl.DB.Select(&jobs, "SELECT * FROM jobs"); err != nil {
		// TODO: handle error properly
		log.Fatal(err)
	}

	ctx.HTML(200, "index", gin.H{
		"jobs":   jobs,
		"noJobs": len(jobs) == 0,
	})
}

func (ctrl *Controller) NewJob(ctx *gin.Context) {
	ctx.HTML(200, "new", gin.H{})
}

func (ctrl *Controller) CreateJob(ctx *gin.Context) {
	var newJobInput NewJob
	ctx.Bind(&newJobInput)

	if errs, valid := newJobInput.validate(); !valid {
		// TODO: send back with errors in session data somehow
		log.Fatal(errs)
	}

	// save job to DB
	if _, err := newJobInput.saveToDB(ctrl.DB); err != nil {
		fmt.Println("failed to save!")
		log.Fatal(err) // TODO: actually handle
	}

	// TODO: send email with edit link
	// TODO: success flash

	ctx.Redirect(302, "/")
}
