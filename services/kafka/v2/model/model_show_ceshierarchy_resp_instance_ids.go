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

type ShowCeshierarchyRespInstanceIds struct {
	// 实例ID。
	Name *string `json:"name,omitempty"`
}

func (o ShowCeshierarchyRespInstanceIds) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowCeshierarchyRespInstanceIds", string(data)}, " ")
}
