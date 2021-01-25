package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateMbTasksReportRequest struct {
	Body *MbTasksReportReq `json:"body,omitempty"`
}

func (o CreateMbTasksReportRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateMbTasksReportRequest struct{}"
	}

	return strings.Join([]string{"CreateMbTasksReportRequest", string(data)}, " ")
}
