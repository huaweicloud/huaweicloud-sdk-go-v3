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

// Response Object
type ChangeApplicationConfigurationResponse struct {
	// 应用ID。
	ApplicationId *string `json:"application_id,omitempty"`
	// 环境ID。
	EnvironmentId *string                             `json:"environment_id,omitempty"`
	Configuration *ApplicationListConfigConfiguration `json:"configuration,omitempty"`
}

func (o ChangeApplicationConfigurationResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ChangeApplicationConfigurationResponse", string(data)}, " ")
}
