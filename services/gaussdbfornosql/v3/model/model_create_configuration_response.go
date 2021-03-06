package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateConfigurationResponse struct {
	Configuration  *CreateConfigurationResult `json:"configuration,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o CreateConfigurationResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateConfigurationResponse struct{}"
	}

	return strings.Join([]string{"CreateConfigurationResponse", string(data)}, " ")
}
