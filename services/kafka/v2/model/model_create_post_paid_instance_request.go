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
type CreatePostPaidInstanceRequest struct {
	ProjectId string                     `json:"project_id"`
	Body      *CreatePostPaidInstanceReq `json:"body,omitempty"`
}

func (o CreatePostPaidInstanceRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreatePostPaidInstanceRequest", string(data)}, " ")
}
