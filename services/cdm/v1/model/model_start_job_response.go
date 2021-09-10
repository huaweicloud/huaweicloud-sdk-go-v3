package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type StartJobResponse struct {
	// 作业运行信息，请参见submission参数说明

	Submissions    *[]StartJobSubmission `json:"submissions,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o StartJobResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "StartJobResponse struct{}"
	}

	return strings.Join([]string{"StartJobResponse", string(data)}, " ")
}
