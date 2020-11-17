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
type ListPluginsRequest struct {
	ProjectId  string `json:"project_id"`
	InstanceId string `json:"instance_id"`
}

func (o ListPluginsRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListPluginsRequest", string(data)}, " ")
}
