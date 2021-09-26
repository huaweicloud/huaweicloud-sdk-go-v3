package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListEditingJobResponse struct {
	Body           *[]QueryEditingJobRsp `json:"body,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListEditingJobResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListEditingJobResponse struct{}"
	}

	return strings.Join([]string{"ListEditingJobResponse", string(data)}, " ")
}
