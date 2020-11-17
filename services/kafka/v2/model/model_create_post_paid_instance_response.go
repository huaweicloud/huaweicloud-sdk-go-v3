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

// Response Object
type CreatePostPaidInstanceResponse struct {
	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`
}

func (o CreatePostPaidInstanceResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreatePostPaidInstanceResponse", string(data)}, " ")
}
