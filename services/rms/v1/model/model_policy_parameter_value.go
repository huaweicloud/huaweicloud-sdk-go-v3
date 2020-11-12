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

// 规则参数值
type PolicyParameterValue struct {
	// 规则参数值
	Value *interface{} `json:"value,omitempty"`
}

func (o PolicyParameterValue) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"PolicyParameterValue", string(data)}, " ")
}
