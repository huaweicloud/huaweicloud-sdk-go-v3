package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowJobDetailResponse struct {
	Job            *JobDetail `json:"job,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ShowJobDetailResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowJobDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowJobDetailResponse", string(data)}, " ")
}
