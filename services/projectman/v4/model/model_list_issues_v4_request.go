package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListIssuesV4Request struct {
	// 项目id

	ProjectId string `json:"project_id"`

	Body *ListIssueRequestV4 `json:"body,omitempty"`
}

func (o ListIssuesV4Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListIssuesV4Request struct{}"
	}

	return strings.Join([]string{"ListIssuesV4Request", string(data)}, " ")
}
