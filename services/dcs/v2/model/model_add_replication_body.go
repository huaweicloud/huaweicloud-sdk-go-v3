/*
 * DCS
 *
 * DCS V2版本API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// 添加副本请求体
type AddReplicationBody struct {
	// 表示指定副本所在的可用区编码。 可用区编码可通过[查询可用区信息](https://support.huaweicloud.com/api-dcs/ListAvailableZones.html)接口查询，可用区必须是有资源的，否则添加失败。
	AzCode *string `json:"az_code,omitempty"`
}

func (o AddReplicationBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AddReplicationBody struct{}"
	}

	return strings.Join([]string{"AddReplicationBody", string(data)}, " ")
}
