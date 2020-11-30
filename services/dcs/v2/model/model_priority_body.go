/*
 * DCS
 *
 * DCS V2版本API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

type PriorityBody struct {
	// 副本优先级，取值范围是0到100，0为默认禁止倒换。
	SlavePriorityWeight int32 `json:"slave_priority_weight"`
}

func (o PriorityBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"PriorityBody", string(data)}, " ")
}
