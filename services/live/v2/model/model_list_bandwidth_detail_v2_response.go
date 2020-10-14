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
type ListBandwidthDetailV2Response struct {
	// 采样数据列表
	DataList   *[]V2BandwidthData `json:"data_list,omitempty"`
	XRequestId *string            `json:"X-request-id,omitempty"`
}

func (o ListBandwidthDetailV2Response) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListBandwidthDetailV2Response", string(data)}, " ")
}
