package main

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"net/url"
)

func signatureForJob(job Job, secret string) string {
	input := fmt.Sprintf(
		"%s:%s:%s:%s",
		job.ID,
		job.Email,
		job.PublishedAt.String(),
		secret,
	)

	hash := sha1.New()
	hash.Write([]byte(input))

	return string(base64.URLEncoding.EncodeToString(hash.Sum(nil)))
}

func signedJobRoute(job Job, config Config) string {
	return fmt.Sprintf(
		"%s/jobs/%s/edit?token=%s",
		config.URL,
		job.ID,
		url.QueryEscape(signatureForJob(job, config.AppSecret)),
	)
}
