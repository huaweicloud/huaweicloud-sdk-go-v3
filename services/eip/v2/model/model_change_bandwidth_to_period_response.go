package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeBandwidthToPeriodResponse Response Object
type ChangeBandwidthToPeriodResponse struct {

	// 转包带宽列表
	BandwidthIds *[]string `json:"bandwidth_ids,omitempty"`

	// 订单ID
	OrderId *string `json:"order_id,omitempty"`

	// 请求ID
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ChangeBandwidthToPeriodResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeBandwidthToPeriodResponse struct{}"
	}

	return strings.Join([]string{"ChangeBandwidthToPeriodResponse", string(data)}, " ")
}
