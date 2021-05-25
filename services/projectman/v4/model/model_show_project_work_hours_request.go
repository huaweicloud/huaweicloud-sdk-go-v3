package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowProjectWorkHoursRequest struct {
	// 项目id

	ProjectId string `json:"project_id"`

	Body *ShowProjectWorkHoursRequestBody `json:"body,omitempty"`
}

func (o ShowProjectWorkHoursRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowProjectWorkHoursRequest struct{}"
	}

	return strings.Join([]string{"ShowProjectWorkHoursRequest", string(data)}, " ")
}
