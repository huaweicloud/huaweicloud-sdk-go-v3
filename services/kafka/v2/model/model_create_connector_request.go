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
type CreateConnectorRequest struct {
	ProjectId  string              `json:"project_id"`
	InstanceId string              `json:"instance_id"`
	Body       *CreateConnectorReq `json:"body,omitempty"`
}

func (o CreateConnectorRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateConnectorRequest", string(data)}, " ")
}
