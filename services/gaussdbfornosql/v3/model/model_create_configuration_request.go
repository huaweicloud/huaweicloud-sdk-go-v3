package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateConfigurationRequest struct {
	Body *CreateConfigurationRequestBody `json:"body,omitempty"`
}

func (o CreateConfigurationRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateConfigurationRequest struct{}"
	}

	return strings.Join([]string{"CreateConfigurationRequest", string(data)}, " ")
}
