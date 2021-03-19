package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateConfigurationRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	ConfigId string `json:"config_id"`

	Body *ConfigurationForUpdate `json:"body,omitempty"`
}

func (o UpdateConfigurationRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateConfigurationRequest struct{}"
	}

	return strings.Join([]string{"UpdateConfigurationRequest", string(data)}, " ")
}
