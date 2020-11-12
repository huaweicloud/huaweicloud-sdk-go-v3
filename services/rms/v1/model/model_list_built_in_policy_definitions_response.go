/*
 * RMS
 *
 * Resource Manager Api
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListBuiltInPolicyDefinitionsResponse struct {
	// 策略定义列表
	Value    *[]PolicyDefinition `json:"value,omitempty"`
	PageInfo *PageInfo           `json:"page_info,omitempty"`
}

func (o ListBuiltInPolicyDefinitionsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListBuiltInPolicyDefinitionsResponse", string(data)}, " ")
}
