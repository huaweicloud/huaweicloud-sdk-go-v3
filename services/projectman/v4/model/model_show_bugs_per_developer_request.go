package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowBugsPerDeveloperRequest struct {
	ProjectId string          `json:"project_id"`
	Body      *MetricRequest2 `json:"body,omitempty"`
}

func (o ShowBugsPerDeveloperRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowBugsPerDeveloperRequest struct{}"
	}

	return strings.Join([]string{"ShowBugsPerDeveloperRequest", string(data)}, " ")
}
