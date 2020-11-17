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

// Request Object
type ShowBackgroundTaskRequest struct {
	ProjectId  string `json:"project_id"`
	InstanceId string `json:"instance_id"`
	TaskId     string `json:"task_id"`
}

func (o ShowBackgroundTaskRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowBackgroundTaskRequest", string(data)}, " ")
}
