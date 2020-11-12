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

// 标签对象
type TagEntity struct {
	// 标签key
	TagKey *string `json:"tag_key,omitempty"`
	// 标签值列表
	TagValue *[]string `json:"tag_value,omitempty"`
}

func (o TagEntity) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"TagEntity", string(data)}, " ")
}
