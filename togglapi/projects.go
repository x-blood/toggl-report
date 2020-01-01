package togglapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"toggl-report/common"
	"toggl-report/togglapimodel"
)

func GetProjectName(projectID uint64) (string, error) {
	client := &http.Client{Timeout: time.Duration(10) * time.Second}
	url := fmt.Sprintf("https://www.toggl.com/api/v8/projects/%d", projectID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.SetBasicAuth(common.GetApiToken(), "api_token")
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var project togglapimodel.Project
	err = json.Unmarshal(getContent(resp), &project)
	if err != nil {
		return "", err
	}
	return project.Data.Name, nil
}

func getContent(resp *http.Response) []byte {
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return b
}
