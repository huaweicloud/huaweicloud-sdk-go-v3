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
type ListInstancesResponse struct {
	// 实例列表
	Instances *[]ListInstancesRespInstances `json:"instances,omitempty"`
	// 实例数量。
	InstanceNum *int32 `json:"instance_num,omitempty"`
}

func (o ListInstancesResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListInstancesResponse", string(data)}, " ")
}
