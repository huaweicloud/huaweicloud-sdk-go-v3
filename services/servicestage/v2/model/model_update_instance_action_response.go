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
type UpdateInstanceActionResponse struct {
	// Job ID。
	JobId *string `json:"job_id,omitempty"`
}

func (o UpdateInstanceActionResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateInstanceActionResponse", string(data)}, " ")
}
