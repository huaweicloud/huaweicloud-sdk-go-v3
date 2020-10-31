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
type ShowJobDetailResponse struct {
	// 部署任务数量。
	TaskCount *int32   `json:"task_count,omitempty"`
	Job       *JobInfo `json:"job,omitempty"`
	// 部署任务列表。
	Tasks *[]TaskInfo `json:"tasks,omitempty"`
}

func (o ShowJobDetailResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowJobDetailResponse", string(data)}, " ")
}
