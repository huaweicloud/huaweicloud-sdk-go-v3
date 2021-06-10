package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListJobInfoDetailResponse struct {
	Jobs           *GetTaskDetailListRspJobs `json:"jobs,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListJobInfoDetailResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListJobInfoDetailResponse struct{}"
	}

	return strings.Join([]string{"ListJobInfoDetailResponse", string(data)}, " ")
}
