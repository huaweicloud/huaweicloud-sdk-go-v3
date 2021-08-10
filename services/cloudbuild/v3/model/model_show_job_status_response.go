package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowJobStatusResponse struct {
	// 任务运行结果

	Result         *bool `json:"result,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowJobStatusResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowJobStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowJobStatusResponse", string(data)}, " ")
}
