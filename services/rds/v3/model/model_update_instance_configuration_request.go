package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateInstanceConfigurationRequest struct {
	XLanguage  *string                                 `json:"X-Language,omitempty"`
	InstanceId string                                  `json:"instance_id"`
	Body       *UpdateInstanceConfigurationRequestBody `json:"body,omitempty"`
}

func (o UpdateInstanceConfigurationRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateInstanceConfigurationRequest struct{}"
	}

	return strings.Join([]string{"UpdateInstanceConfigurationRequest", string(data)}, " ")
}
