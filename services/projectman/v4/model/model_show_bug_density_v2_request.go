package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowBugDensityV2Request struct {
	ProjectId string           `json:"project_id"`
	Body      *MetricRequestV2 `json:"body,omitempty"`
}

func (o ShowBugDensityV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowBugDensityV2Request struct{}"
	}

	return strings.Join([]string{"ShowBugDensityV2Request", string(data)}, " ")
}
