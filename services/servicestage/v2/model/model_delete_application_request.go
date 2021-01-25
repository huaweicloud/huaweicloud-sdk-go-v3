package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteApplicationRequest struct {
	ApplicationId string `json:"application_id"`
}

func (o DeleteApplicationRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteApplicationRequest struct{}"
	}

	return strings.Join([]string{"DeleteApplicationRequest", string(data)}, " ")
}
