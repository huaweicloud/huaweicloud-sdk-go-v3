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
type CreateInstanceResponse struct {
	// 应用组件实例ID。
	InstanceId *string `json:"instance_id,omitempty"`
	// Job ID，用于查询创建任务信息。
	JobId *string `json:"job_id,omitempty"`
}

func (o CreateInstanceResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateInstanceResponse", string(data)}, " ")
}
