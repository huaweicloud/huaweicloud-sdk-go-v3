package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateDefaultConfigRequest struct {
	Body *DdosConfig `json:"body,omitempty"`
}

func (o CreateDefaultConfigRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateDefaultConfigRequest struct{}"
	}

	return strings.Join([]string{"CreateDefaultConfigRequest", string(data)}, " ")
}
