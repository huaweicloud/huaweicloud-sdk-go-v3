package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowIssueCompletionRateRequest struct {
	ProjectId string `json:"project_id"`
}

func (o ShowIssueCompletionRateRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowIssueCompletionRateRequest struct{}"
	}

	return strings.Join([]string{"ShowIssueCompletionRateRequest", string(data)}, " ")
}
