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
type ChangeEnvironmentRequest struct {
	EnvironmentId string             `json:"environment_id"`
	Body          *EnvironmentModify `json:"body,omitempty"`
}

func (o ChangeEnvironmentRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ChangeEnvironmentRequest", string(data)}, " ")
}
