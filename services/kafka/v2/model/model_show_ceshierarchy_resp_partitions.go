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

type ShowCeshierarchyRespPartitions struct {
	// 分区名称。
	Name *string `json:"name,omitempty"`
}

func (o ShowCeshierarchyRespPartitions) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowCeshierarchyRespPartitions", string(data)}, " ")
}
