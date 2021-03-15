package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowCompletionRateRequest struct {
	ProjectId string          `json:"project_id"`
	Body      *MetricRequest3 `json:"body,omitempty"`
}

func (o ShowCompletionRateRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowCompletionRateRequest struct{}"
	}

	return strings.Join([]string{"ShowCompletionRateRequest", string(data)}, " ")
}
