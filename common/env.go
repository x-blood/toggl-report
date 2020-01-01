package common

import (
	"fmt"
	"os"
	"time"
)

const DateFormat = "2006-01-02"

func GetApiToken() string {
	return os.Getenv("TOGGL_API_TOKEN")
}

func GetSlackWebHookURL() string {
	return os.Getenv("SLACK_WEB_HOOK_URL")
}

func GetSlackChannel() string {
	return os.Getenv("SLACK_CHANNEL_NAME")
}

func GetTogglWorkspaceID() string {
	return os.Getenv("TOGGL_WORKSPACE_ID")
}

func GetStartDate() string {
	return fmt.Sprintf("%sT00%%3A00%%3A00%%2B09%%3A00",
		time.Now().In(JST()).AddDate(0, 0, -1).Format(DateFormat))
}

func GetEndDate() string {
	return fmt.Sprintf("%sT23%%3A59%%3A59%%2B09%%3A00",
		time.Now().In(JST()).AddDate(0, 0, -1).Format(DateFormat))
}

func GetTargetDate() string {
	return time.Now().In(JST()).AddDate(0, 0, -1).Format(DateFormat)
}

func JST() *time.Location {
	l, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		fmt.Println(err.Error())
	}
	return l
}
