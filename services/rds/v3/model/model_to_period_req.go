package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ToPeriodReq RDS实例按需转包周期请求
type ToPeriodReq struct {

	// 周期类型。MONTH：月；YEAR：年
	PeriodType string `json:"period_type"`

	// 周期数。
	PeriodNum int32 `json:"period_num"`

	// 是否自动支付。YES：自动扣费；NO：手动支付（默认）
	AutoPayPolicy *string `json:"auto_pay_policy,omitempty"`

	// 是否到期自动续期。YES：自动续费；NO：不自动续费（默认）
	AutoRenewPolicy *string `json:"auto_renew_policy,omitempty"`
}

func (o ToPeriodReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ToPeriodReq struct{}"
	}

	return strings.Join([]string{"ToPeriodReq", string(data)}, " ")
}
