package main

import (
	"fmt"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func postToTwitter(job Job, config Config) error {
	tweetStr := tweetFromJob(job, config)

	oa := oauth1.NewConfig(config.Twitter.APIKey, config.Twitter.APISecretKey)
	token := oauth1.NewToken(
		config.Twitter.AccessToken,
		config.Twitter.AccessTokenSecret,
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

func tweetFromJob(job Job, config Config) string {
	return fmt.Sprintf(
		"A job was posted! -- %s at %s\n\nMore info at %s/jobs/%s",
		job.Position,
		job.Organization,
		config.URL,
		job.ID,
	)
}
