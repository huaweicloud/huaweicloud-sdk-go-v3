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

type ShowGroupsRespGroupAssignment struct {
	// topic名称。
	Topic *string `json:"topic,omitempty"`
	// 分区列表。
	Partitions *[]string `json:"partitions,omitempty"`
}

func (o ShowGroupsRespGroupAssignment) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowGroupsRespGroupAssignment", string(data)}, " ")
}
