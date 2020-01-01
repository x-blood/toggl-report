package domain

import (
	"fmt"
	"strings"
	"toggl-report/togglapi"
)

type SlackFormat struct {
	Project []Project
}

type Project struct {
	ID           uint64
	Name         string
	Duration     uint64
	Descriptions []Description
}

type Description struct {
	Name     string
	Duration uint64
}

func (s SlackFormat) GetTotalDuration() uint64 {
	var result uint64
	for _, project := range s.Project {
		result = result + project.GetProjectDuration()
	}

	return result
}

func (s SlackFormat) GetTotalDurationText() string {
	var result uint64
	for _, project := range s.Project {
		result = result + project.GetProjectDuration()
	}

	hour := result / 3600
	min := (result % 3600) / 60
	sec := result % 60

	return fmt.Sprintf("%2d時間%02d分%02d秒", hour, min, sec)
}

func (p Project) GetProjectDuration() uint64 {
	var result uint64
	for _, description := range p.Descriptions {
		result = result + description.Duration
	}

	return result
}

func (p Project) GetProjectDurationText() string {
	var result uint64
	for _, description := range p.Descriptions {
		result = result + description.Duration
	}

	hour := result / 3600
	min := (result % 3600) / 60
	sec := result % 60

	return fmt.Sprintf("%2d:%02d:%02d", hour, min, sec)
}

func (p Project) GetDescriptionText() string {
	var sb strings.Builder
	descriptionDurationMap := p.GetDescriptionDurationMap()
	for key, value := range descriptionDurationMap {
		hour := value / 3600
		min := (value % 3600) / 60
		sec := value % 60
		sb.WriteString(fmt.Sprintf("• %s : %s\n", key, fmt.Sprintf("%2d:%02d:%02d", hour, min, sec)))
	}
	return sb.String()
}

func (p Project) GetDescriptionDurationMap() map[string]uint64 {
	result := map[string]uint64{}
	for _, description := range p.Descriptions {
		if _, ok := result[description.Name]; !ok {
			result[description.Name] = description.Duration
		} else {
			result[description.Name] = result[description.Name] + description.Duration
		}
	}
	return result
}

func GetSlackFormat(timeEntries []TimeEntry) (*SlackFormat, error) {
	var result []Project
	projects, err := getProjectList(timeEntries)
	if err != nil {
		return nil, err
	}
	for _, project := range projects {
		descriptions := getDescriptionListAsProjectID(project.ID, timeEntries)
		project.Descriptions = descriptions
		result = append(result, project)
	}

	return &SlackFormat{Project: result}, nil
}

func getProjectList(timeEntries []TimeEntry) ([]Project, error) {
	var result []Project
	projectNameMap := map[uint64]string{}
	for _, timeEntry := range timeEntries {
		if _, ok := projectNameMap[timeEntry.ProjectID]; !ok {
			if timeEntry.ProjectID != 0 {
				projectName, err := togglapi.GetProjectName(timeEntry.ProjectID)
				if err != nil {
					return nil, err
				}

				projectNameMap[timeEntry.ProjectID] = projectName
				result = append(result, Project{
					ID:           timeEntry.ProjectID,
					Name:         projectName,
					Duration:     0,
					Descriptions: []Description{},
				})
			} else {
				projectNameMap[timeEntry.ProjectID] = "その他"
				result = append(result, Project{
					ID:           timeEntry.ProjectID,
					Name:         "その他",
					Duration:     0,
					Descriptions: []Description{},
				})
			}
		}
	}
	return result, nil
}

func getDescriptionListAsProjectID(projectID uint64, timeEntries []TimeEntry) []Description {
	var result []Description
	for _, timeEntry := range timeEntries {
		if timeEntry.ProjectID == projectID {
			result = append(result, Description{
				Name:     timeEntry.Description,
				Duration: timeEntry.Duration,
			})
		}
	}
	return result
}
