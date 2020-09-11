/*
 * kps
 *
 * kps v3 版本API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListKeypairTaskRequest struct {
	TaskId string `json:"task_id"`
}

func (o ListKeypairTaskRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListKeypairTaskRequest", string(data)}, " ")
}
