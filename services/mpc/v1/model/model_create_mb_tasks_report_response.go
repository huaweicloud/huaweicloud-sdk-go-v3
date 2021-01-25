package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateMbTasksReportResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateMbTasksReportResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateMbTasksReportResponse struct{}"
	}

	return strings.Join([]string{"CreateMbTasksReportResponse", string(data)}, " ")
}
