package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateChInstanceInfoPayInfo 计费信息。
type CreateChInstanceInfoPayInfo struct {

	// 计费模式。默认0。 取值范围： - 0：按需计费 - 1：包周期
	PayModel string `json:"pay_model"`

	// 包周期计费ID。
	OrderId string `json:"order_id"`

	// 包周期周期。
	Period string `json:"period"`

	// 包周期周期类型。 取值范围： - 2：包月 - 3：包年
	PeriodType string `json:"period_type"`

	// 包周期是否自动续费。 取值范围： - 1：自动续费 - 0：不自动续费
	IsAutoRenew string `json:"is_auto_renew"`
}

func (o CreateChInstanceInfoPayInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateChInstanceInfoPayInfo struct{}"
	}

	return strings.Join([]string{"CreateChInstanceInfoPayInfo", string(data)}, " ")
}
