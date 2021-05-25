package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteIssueV4Request struct {
	// 项目id

	ProjectId string `json:"project_id"`
	// 工作项id

	IssueId int32 `json:"issue_id"`
}

func (o DeleteIssueV4Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteIssueV4Request struct{}"
	}

	return strings.Join([]string{"DeleteIssueV4Request", string(data)}, " ")
}
