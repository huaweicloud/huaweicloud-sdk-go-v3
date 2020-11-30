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
type UpdatePluginsResponse struct {
	// 后台任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdatePluginsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdatePluginsResponse", string(data)}, " ")
}
