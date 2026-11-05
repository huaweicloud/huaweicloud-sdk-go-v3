package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PeriodReq struct {

	// 订购周期类型。 - 2: 包月。 - 3: 包年。
	PeriodType int32 `json:"period_type"`

	// 订购周期数。 - 若选择包月（参数范围：1-9）。 - 若选择包年（参数范围：1-3）。
	PeriodNum int32 `json:"period_num"`

	// 是否自动续订，为空时表示不自动续订 - 1: 自动续订。 - 0: 不自动续订（默认）。
	IsAutoRenew *int32 `json:"is_auto_renew,omitempty"`

	IsAutoPay *int32 `json:"is_auto_pay,omitempty"`

	// 云服务ConsoleURL。 订购订单支付完成后，客户可以通过此URL跳转到云服务Console页面查看信息。（仅手动支付时涉及）。
	ConsoleUrl *string `json:"console_url,omitempty"`
}

func (o PeriodReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PeriodReq struct{}"
	}

	return strings.Join([]string{"PeriodReq", string(data)}, " ")
}
