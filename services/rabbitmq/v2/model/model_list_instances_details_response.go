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

// Response Object
type ListInstancesDetailsResponse struct {
	// 实例列表。
	Instances *[]ListInstancesRespInstances `json:"instances,omitempty"`
	// 实例个数。
	InstanceNum *int32 `json:"instance_num,omitempty"`
}

func (o ListInstancesDetailsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListInstancesDetailsResponse", string(data)}, " ")
}
