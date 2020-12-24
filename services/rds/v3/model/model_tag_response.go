/*
 * RDS
 *
 * API v3
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// 标签信息。
type TagResponse struct {
	// 标签键。
	Key string `json:"key"`
	// 标签值。
	Value string `json:"value"`
}

func (o TagResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"TagResponse", string(data)}, " ")
}
