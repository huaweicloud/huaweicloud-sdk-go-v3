package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StarRocksCreateRequestPayInfo 支付信息。包周期场景必填。
type StarRocksCreateRequestPayInfo struct {

	// 计费模式，默认0。包周期场景必填。 取值范围： - 0：按需计费 - 1：包周期
	PayModel *string `json:"pay_model,omitempty"`

	// 包周期周期。包周期场景必填。
	Period *string `json:"period,omitempty"`

	// 包周期周期类型。包周期场景必填。 取值范围： - 2：包月 - 3：包年
	PeriodType *string `json:"period_type,omitempty"`

	// 包周期是否自动续费。包周期场景必填。 取值范围： - 1：自动续费 - 0：不自动续费
	IsAutoRenew *string `json:"is_auto_renew,omitempty"`
}

func (o StarRocksCreateRequestPayInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StarRocksCreateRequestPayInfo struct{}"
	}

	return strings.Join([]string{"StarRocksCreateRequestPayInfo", string(data)}, " ")
}
