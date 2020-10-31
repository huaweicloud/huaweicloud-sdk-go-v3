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

type InstanceAction struct {
	Action     *InstanceActionType       `json:"action"`
	Parameters *InstanceActionParameters `json:"parameters,omitempty"`
}

func (o InstanceAction) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"InstanceAction", string(data)}, " ")
}
