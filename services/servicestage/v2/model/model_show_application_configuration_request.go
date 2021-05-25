package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowApplicationConfigurationRequest struct {
	// 应用ID。

	ApplicationId string `json:"application_id"`
	// 环境ID，如果未提供，查询所有环境。

	EnvironmentId *string `json:"environment_id,omitempty"`
}

func (o ShowApplicationConfigurationRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowApplicationConfigurationRequest struct{}"
	}

	return strings.Join([]string{"ShowApplicationConfigurationRequest", string(data)}, " ")
}
