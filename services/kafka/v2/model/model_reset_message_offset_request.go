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
type ResetMessageOffsetRequest struct {
	ProjectId  string                 `json:"project_id"`
	InstanceId string                 `json:"instance_id"`
	Group      string                 `json:"group"`
	Body       *ResetMessageOffsetReq `json:"body,omitempty"`
}

func (o ResetMessageOffsetRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ResetMessageOffsetRequest", string(data)}, " ")
}
