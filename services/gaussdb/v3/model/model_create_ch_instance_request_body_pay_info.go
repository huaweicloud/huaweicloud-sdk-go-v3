package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateChInstanceRequestBodyPayInfo 支付信息。
type CreateChInstanceRequestBodyPayInfo struct {

	// 计费模式，默认0。 取值范围： - 0：按需计费 - 1：包周期
	PayModel *string `json:"pay_model,omitempty"`

	// 包周期计费ID。
	OrderId *string `json:"order_id,omitempty"`

	// 包周期周期。
	Period *int32 `json:"period,omitempty"`

	// 包周期周期类型。 取值范围： - 2：包月 - 3：包年
	PeriodType *int32 `json:"period_type,omitempty"`

	// 包周期是否自动续费。 取值范围： - 1：自动续费 - 0：不自动续费
	IsAutoRenew *int32 `json:"is_auto_renew,omitempty"`
}

func (o CreateChInstanceRequestBodyPayInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateChInstanceRequestBodyPayInfo struct{}"
	}

	return strings.Join([]string{"CreateChInstanceRequestBodyPayInfo", string(data)}, " ")
}
