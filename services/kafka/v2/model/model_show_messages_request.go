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
type ShowMessagesRequest struct {
	ProjectId  string  `json:"project_id"`
	InstanceId string  `json:"instance_id"`
	Topic      string  `json:"topic"`
	StartTime  *string `json:"start_time,omitempty"`
	EndTime    *string `json:"end_time,omitempty"`
	Limit      *int32  `json:"limit,omitempty"`
	Offset     *int32  `json:"offset,omitempty"`
	Partition  *string `json:"partition,omitempty"`
}

func (o ShowMessagesRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowMessagesRequest", string(data)}, " ")
}
