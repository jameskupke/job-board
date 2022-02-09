package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type SlackMessage struct {
	Text string `json:"text"`
}

func postToSlack(job Job, config Config) error {
	message := slackMessageFromJob(job, config)
	messageStr, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal slack message: %w", err)
	}

	_, err = http.Post(config.SlackHook, "application/json", bytes.NewReader(messageStr))
	if err != nil {
		return fmt.Errorf("failed to post to slack: %w", err)
	}

	return nil
}

func slackMessageFromJob(job Job, config Config) SlackMessage {
	text := fmt.Sprintf(
		"A new job was posted!\n> *<%s/jobs/%s|%s @ %s>*",
		config.URL,
		job.ID,
		job.Position,
		job.Organization,
	)
	return SlackMessage{Text: text}
}
