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

type DataIpRequest struct {
	// 内网ip
	NewIp string `json:"new_ip"`
}

func (o DataIpRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DataIpRequest", string(data)}, " ")
}
