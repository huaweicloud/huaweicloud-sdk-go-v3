package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowJobStatusResponse struct {
	// 作业运行信息，详见submissions参数说明。

	Submissions    *[]Submission `json:"submissions,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowJobStatusResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowJobStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowJobStatusResponse", string(data)}, " ")
}
