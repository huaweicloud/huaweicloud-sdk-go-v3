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

type FailoverModeRequest struct {
	// 数据库主备同步模式
	Mode string `json:"mode"`
}

func (o FailoverModeRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"FailoverModeRequest", string(data)}, " ")
}
