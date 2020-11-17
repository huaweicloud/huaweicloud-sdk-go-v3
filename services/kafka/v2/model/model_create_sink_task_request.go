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
type CreateSinkTaskRequest struct {
	ProjectId   string             `json:"project_id"`
	ConnectorId string             `json:"connector_id"`
	Body        *CreateSinkTaskReq `json:"body,omitempty"`
}

func (o CreateSinkTaskRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateSinkTaskRequest", string(data)}, " ")
}
