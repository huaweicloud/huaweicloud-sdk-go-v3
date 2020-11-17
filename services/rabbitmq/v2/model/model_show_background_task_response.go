/*
 * RabbitMQ
 *
 * RabbitMQ Document API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowBackgroundTaskResponse struct {
	// 任务数量。
	TaskCount *string `json:"task_count,omitempty"`
	// 任务列表。
	Tasks *[]ListBackgroundTasksRespTasks `json:"tasks,omitempty"`
}

func (o ShowBackgroundTaskResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowBackgroundTaskResponse", string(data)}, " ")
}
