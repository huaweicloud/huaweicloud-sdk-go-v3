/*
 * BMS
 *
 * BMS Open API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// os-start字段数据结构说明
type StartServersInfo struct {
	// 裸金属服务器ID列表，详情请参见表3 servers字段数据结构说明
	Servers []ServersList `json:"servers"`
}

func (o StartServersInfo) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"StartServersInfo", string(data)}, " ")
}
