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
type ShowInstanceDetailRequest struct {
	ApplicationId string `json:"application_id"`
	ComponentId   string `json:"component_id"`
	InstanceId    string `json:"instance_id"`
}

func (o ShowInstanceDetailRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowInstanceDetailRequest", string(data)}, " ")
}
