package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteIaConfigResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteIaConfigResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteIaConfigResponse struct{}"
	}

	return strings.Join([]string{"DeleteIaConfigResponse", string(data)}, " ")
}
