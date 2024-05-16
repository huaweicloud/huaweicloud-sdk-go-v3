package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SrCreateInstanceRspInstancePayInfo 计费信息。
type SrCreateInstanceRspInstancePayInfo struct {

	// 计费模式。 - 0：按需计费 - 1：包周期  StarRocks实例当前只支持按需计费，默认值为0
	PayModel *string `json:"pay_model,omitempty"`

	// 包周期计费ID。
	OrderId *string `json:"order_id,omitempty"`

	// 包周期周期。
	Period *string `json:"period,omitempty"`

	// 包周期周期类型。
	PeriodType *string `json:"period_type,omitempty"`

	// 包周期是否自动续费。
	IsAutoRenew *string `json:"is_auto_renew,omitempty"`
}

func (o SrCreateInstanceRspInstancePayInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SrCreateInstanceRspInstancePayInfo struct{}"
	}

	return strings.Join([]string{"SrCreateInstanceRspInstancePayInfo", string(data)}, " ")
}
