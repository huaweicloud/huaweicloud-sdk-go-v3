/*
 * DCS
 *
 * DCS V2版本API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteBackgroundTaskRequest struct {
	InstanceId string `json:"instance_id"`
	TaskId     string `json:"task_id"`
}

func (o DeleteBackgroundTaskRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteBackgroundTaskRequest", string(data)}, " ")
}
