package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdatePrePaidBandwidthResponse struct {
	Bandwidth *BandwidthResp `json:"bandwidth,omitempty" xml:"bandwidth"`

	// 订单号（包周期场景返回该字段）
	OrderId        *string `json:"order_id,omitempty" xml:"order_id"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdatePrePaidBandwidthResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePrePaidBandwidthResponse struct{}"
	}

	return strings.Join([]string{"UpdatePrePaidBandwidthResponse", string(data)}, " ")
}
