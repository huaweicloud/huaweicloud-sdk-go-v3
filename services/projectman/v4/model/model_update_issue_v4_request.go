package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateIssueV4Request struct {
	// 项目id

	ProjectId string `json:"project_id"`
	// 工作项id

	IssueId int32 `json:"issue_id"`

	Body *IssueRequestV4 `json:"body,omitempty"`
}

func (o UpdateIssueV4Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateIssueV4Request struct{}"
	}

	return strings.Join([]string{"UpdateIssueV4Request", string(data)}, " ")
}
