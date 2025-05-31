package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
	"toggl-report/common"
	"toggl-report/domain"
	"toggl-report/infrastructure"
)

func Do() error {
	client := &http.Client{Timeout: time.Duration(10) * time.Second}
	url := fmt.Sprintf("https://api.track.toggl.com/api/v9/me/time_entries?start_date=%s&end_date=%s", common.GetStartDate(), common.GetEndDate())
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
		return err
	}
	req.SetBasicAuth(common.GetApiToken(), "api_token")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer resp.Body.Close()

	var timeEntries []domain.TimeEntry
	err = json.Unmarshal(getContent(resp), &timeEntries)
	if err != nil {
		log.Fatal(err)
	}

	// ワークスペースIDを取得して、uint64型に変換
	workspaceIDStr := common.GetTogglWorkspaceID()
	workspaceID, err := strconv.ParseUint(workspaceIDStr, 10, 64)
	if err != nil {
		log.Fatal(err)
		return err
	}

	// 指定されたワークスペースIDに一致するTimeEntryだけをフィルタリング
	filteredTimeEntries := domain.FilterByWorkspaceID(timeEntries, workspaceID)

	slackFormat, err := domain.GetSlackFormat(filteredTimeEntries)
	if err != nil {
		log.Fatal(err)
		return err
	}

	if slackFormat.GetTotalDuration() != 0 {
		err = infrastructure.NotifyToSlack(*slackFormat)
		if err != nil {
			log.Fatal(err)
			return err
		}
	}
	return nil
}

func getContent(resp *http.Response) []byte {
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return b
}
