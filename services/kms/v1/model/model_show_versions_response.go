/*
 * kms
 *
 * KMS v1.0 API, open API
 *
 */

package model

import (
	"encoding/json"
	"strings"
)

// Response Object
type ShowVersionsResponse struct {
	// 描述version 相关对象的列表，详情请参见 versions字段数据结构说明。
	Versions []ApiVersionDetail `json:"versions,omitempty"`
}

func (o ShowVersionsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowVersionsResponse", string(data)}, " ")
}
