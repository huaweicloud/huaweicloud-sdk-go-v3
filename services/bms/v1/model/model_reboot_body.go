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

// 重启裸金属服务器接口请求结构体
type RebootBody struct {
	Reboot *ServersInfoType `json:"reboot"`
}

func (o RebootBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"RebootBody", string(data)}, " ")
}
