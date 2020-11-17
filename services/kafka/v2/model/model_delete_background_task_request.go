/*
 * Kafka
 *
 * Kafka Document API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteBackgroundTaskRequest struct {
	ProjectId  string `json:"project_id"`
	InstanceId string `json:"instance_id"`
	TaskId     string `json:"task_id"`
}

func (o DeleteBackgroundTaskRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteBackgroundTaskRequest", string(data)}, " ")
}
