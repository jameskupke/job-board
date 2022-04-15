package services

import (
	"fmt"

	"github.com/devict/job-board/pkg/config"
	"github.com/devict/job-board/pkg/data"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func PostToTwitter(job data.Job, c config.Config) error {
	tweetStr := tweetFromJob(job, c)

	oa := oauth1.NewConfig(c.Twitter.APIKey, c.Twitter.APISecretKey)
	token := oauth1.NewToken(
		c.Twitter.AccessToken,
		c.Twitter.AccessTokenSecret,
	)
	httpClient := oa.Client(oauth1.NoContext, token)

	twClient := twitter.NewClient(httpClient)

	// TODO: check for failures in the resp object?
	_, _, err := twClient.Statuses.Update(tweetStr, nil)
	if err != nil {
		return fmt.Errorf("failed to post to twitter: %w", err)
	}

	return nil
}

func tweetFromJob(job data.Job, c config.Config) string {
	return fmt.Sprintf(
		"A job was posted! -- %s at %s\n\nMore info at %s/jobs/%s",
		job.Position,
		job.Organization,
		c.URL,
		job.ID,
	)
}
