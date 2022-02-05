package main

import (
	"fmt"
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type Controller struct {
	DB *sqlx.DB
}

func (ctrl *Controller) Index(ctx *gin.Context) {
	jobs, err := getAllJobs(ctrl.DB)
	if err != nil {
		panic(err) // TODO: handle this
	}

	ctx.HTML(200, "index", addFlash(ctx, gin.H{
		"jobs":   jobs,
		"noJobs": len(jobs) == 0,
	}))
}

func (ctrl *Controller) NewJob(ctx *gin.Context) {
	session := sessions.Default(ctx)

	fields := []string{"position", "organization", "url", "description", "email"}

	tVars := gin.H{}
	for _, k := range fields {
		tVars[k] = session.Flashes(k + "_err")
	}

	ctx.HTML(200, "new", addFlash(ctx, tVars))
}

func (ctrl *Controller) CreateJob(ctx *gin.Context) {
	var newJobInput NewJob
	ctx.Bind(&newJobInput)

	session := sessions.Default(ctx)

	if errs := newJobInput.validate(); len(errs) != 0 {
		for k, v := range errs {
			session.AddFlash(v, k+"_err")
		}
		session.Save()

		ctx.Redirect(302, "/new")
		return
	}

	if _, err := newJobInput.saveToDB(ctrl.DB); err != nil {
		log.Print(fmt.Errorf("failed to save job to db: %w", err))

		session.AddFlash("Error creating job")
		session.Save()

		ctx.Redirect(302, "/new")
		return
	}

	// TODO: send email with edit link

	session.AddFlash("Job created!")
	session.Save()

	ctx.Redirect(302, "/")
}

func (ctrl *Controller) ViewJob(ctx *gin.Context) {
	id := ctx.Param("id")
	job, err := getJob(id, ctrl.DB)
	if err != nil {
		panic(err) // TODO: err
	}
	ctx.HTML(200, "view", gin.H{"job": job})
}

func addFlash(ctx *gin.Context, base gin.H) gin.H {
	session := sessions.Default(ctx)
	base["flashes"] = session.Flashes()
	return base
}
