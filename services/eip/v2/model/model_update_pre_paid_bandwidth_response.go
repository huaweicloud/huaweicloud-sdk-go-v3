/*
 * EIP
 *
 * 云服务接口
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdatePrePaidBandwidthResponse struct {
	Bandwidth *BandwidthResp `json:"bandwidth,omitempty"`
	// 订单号（包周期场景返回该字段）
	OrderId *string `json:"order_id,omitempty"`
}

func (o UpdatePrePaidBandwidthResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdatePrePaidBandwidthResponse", string(data)}, " ")
}
