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
type ShowUpBandwidthV2Response struct {
	// 采样数据列表
	DataList   *[]V2BandwidthData `json:"data_list,omitempty"`
	XRequestId *string            `json:"X-request-id,omitempty"`
}

func (o ShowUpBandwidthV2Response) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowUpBandwidthV2Response", string(data)}, " ")
}
