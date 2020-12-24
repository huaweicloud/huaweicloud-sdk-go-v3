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

type UpdateDbPortRequest struct {
	// MySQL端口号范围：大于等于1024，小于等于65535，不包含12017和33071。
	Port int32 `json:"port"`
}

func (o UpdateDbPortRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateDbPortRequest", string(data)}, " ")
}
