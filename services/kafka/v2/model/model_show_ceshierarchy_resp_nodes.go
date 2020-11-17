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

type ShowCeshierarchyRespNodes struct {
	// 节点名称。
	Name *string `json:"name,omitempty"`
}

func (o ShowCeshierarchyRespNodes) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowCeshierarchyRespNodes", string(data)}, " ")
}
