package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowFileInfoResponse struct {
	Jobs *[]Job `json:"jobs,omitempty"`

	Scripts        *[]Script `json:"scripts,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowFileInfoResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowFileInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowFileInfoResponse", string(data)}, " ")
}
