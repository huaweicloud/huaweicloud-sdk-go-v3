package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowMysqlJobInfoResponse struct {
	Job            *GetJobInfoDetail `json:"job,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowMysqlJobInfoResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowMysqlJobInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowMysqlJobInfoResponse", string(data)}, " ")
}
