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

type FailoverStrategyRequest struct {
	// 可用性策略
	RepairStrategy string `json:"repairStrategy"`
}

func (o FailoverStrategyRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"FailoverStrategyRequest", string(data)}, " ")
}
