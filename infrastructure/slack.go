package infrastructure

import (
	"fmt"
	"toggl-report/common"
	"toggl-report/domain"

	"github.com/ashwanthkumar/slack-go-webhook"
)

const (
	USERNAME = "Daily Report"
)

func NotifyToSlack(slackFormat domain.SlackFormat) error {
	return postSlack(slackFormat)
}

func postSlack(slackFormat domain.SlackFormat) error {
	attachment := slack.Attachment{}
	for _, project := range slackFormat.Project {
		field := slack.Field{Title: fmt.Sprintf("%s : %s", project.Name, project.GetProjectDurationText()), Value: project.GetDescriptionText()}
		attachment.AddField(field)
	}
	color := "good"
	attachment.Color = &color
	payload := slack.Payload{
		IconEmoji:   ":city_sunrise:",
		Username:    USERNAME,
		Channel:     common.GetSlackChannel(),
		Text:        fmt.Sprintf("%sのレポートです。チャージ時間は%sです。", common.GetTargetDate(), slackFormat.GetTotalDurationText()),
		Attachments: []slack.Attachment{attachment},
	}
	err := slack.Send(common.GetSlackWebHookURL(), "", payload)
	if err != nil {
		return err[0]
	}
	return nil
}
