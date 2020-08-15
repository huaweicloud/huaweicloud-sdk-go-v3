/*
 * LiveAPI
 *
 * 直播服务源站所有接口
 *
 */

package model

import (
	"encoding/json"
	"strings"
)

// Response Object
type ShowBandwidthResponse struct {
	// 查询结果的总元素数量
	Total int32 `json:"total,omitempty"`
	// 带宽信息
	BandwidthInfo []BandwidthInfo `json:"bandwidth_info,omitempty"`
}

func (o ShowBandwidthResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowBandwidthResponse", string(data)}, " ")
}
