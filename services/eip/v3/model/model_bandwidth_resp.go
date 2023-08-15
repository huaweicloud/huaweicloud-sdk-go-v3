package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BandwidthResp 弹性公网IP中的bandwidth对象，存储公网IP绑定的带宽信息
type BandwidthResp struct {

	// - 功能说明：带宽ID
	Id *string `json:"id,omitempty"`

	// - 功能说明：带宽大小
	Size *int32 `json:"size,omitempty"`

	// - 功能说明：类型  \"WHOLE\"为共享带宽，\"PER\"为独占带宽
	ShareType *string `json:"share_type,omitempty"`

	// - 功能说明：带宽计费模式
	ChargeMode *string `json:"charge_mode,omitempty"`

	// - 功能说明：带宽名称
	Name *string `json:"name,omitempty"`

	// - 功能说明：带宽的订单信息
	BillingInfo *string `json:"billing_info,omitempty"`
}

func (o BandwidthResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BandwidthResp struct{}"
	}

	return strings.Join([]string{"BandwidthResp", string(data)}, " ")
}
