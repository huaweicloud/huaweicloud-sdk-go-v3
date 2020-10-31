/*
 * ServiceStage
 *
 * ServiceStage的API,包括应用管理和仓库授权管理
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowApplicationConfigurationRequest struct {
	ApplicationId string  `json:"application_id"`
	EnvironmentId *string `json:"environment_id,omitempty"`
}

func (o ShowApplicationConfigurationRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowApplicationConfigurationRequest", string(data)}, " ")
}
