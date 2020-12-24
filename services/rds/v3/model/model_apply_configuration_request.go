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

type ApplyConfigurationRequest struct {
	// 实例ID列表。
	InstanceIds []string `json:"instance_ids"`
}

func (o ApplyConfigurationRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ApplyConfigurationRequest", string(data)}, " ")
}
