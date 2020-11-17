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
type ResizeInstanceResponse struct {
	// 规格变更任务ID。
	JobId *string `json:"job_id,omitempty"`
}

func (o ResizeInstanceResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ResizeInstanceResponse", string(data)}, " ")
}
