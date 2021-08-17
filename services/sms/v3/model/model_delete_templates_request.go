package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteTemplatesRequest struct {
	Body *DeletetemplatesReq `json:"body,omitempty"`
}

func (o DeleteTemplatesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteTemplatesRequest struct{}"
	}

	return strings.Join([]string{"DeleteTemplatesRequest", string(data)}, " ")
}
