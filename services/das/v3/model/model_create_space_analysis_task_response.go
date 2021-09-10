package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateSpaceAnalysisTaskResponse struct {
	// 执行时间，毫秒为单位的时间戳

	ExecutionTime  *int64 `json:"execution_time,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateSpaceAnalysisTaskResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateSpaceAnalysisTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateSpaceAnalysisTaskResponse", string(data)}, " ")
}
