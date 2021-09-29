package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowGaussMySqlJobInfoResponse struct {
	Job            *GetJobInfoDetail `json:"job,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowGaussMySqlJobInfoResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowGaussMySqlJobInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowGaussMySqlJobInfoResponse", string(data)}, " ")
}
