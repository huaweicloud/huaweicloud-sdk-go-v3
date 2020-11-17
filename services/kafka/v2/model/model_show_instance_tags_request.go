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
type ShowInstanceTagsRequest struct {
	ProjectId  string `json:"project_id"`
	InstanceId string `json:"instance_id"`
}

func (o ShowInstanceTagsRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowInstanceTagsRequest", string(data)}, " ")
}
