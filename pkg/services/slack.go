package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/devict/job-board/pkg/config"
	"github.com/devict/job-board/pkg/data"
)

type SlackMessage struct {
	Text string `json:"text"`
}

func PostToSlack(job data.Job, c config.Config) error {
	message := slackMessageFromJob(job, c)
	messageStr, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal slack message: %w", err)
	}

	_, err = http.Post(c.SlackHook, "application/json", bytes.NewReader(messageStr))
	if err != nil {
		return fmt.Errorf("failed to post to slack: %w", err)
	}

	return nil
}

func slackMessageFromJob(job data.Job, c config.Config) SlackMessage {
	text := fmt.Sprintf(
		"A new job was posted!\n> *<%s/jobs/%s|%s @ %s>*",
		c.URL,
		job.ID,
		job.Position,
		job.Organization,
	)
	return SlackMessage{Text: text}
}
