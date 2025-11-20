package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PostPaidChargeInfos 支付信息
type PostPaidChargeInfos struct {

	// - 参数解释：付费方式（后付费，即按需付费）。 - 约束限制：不涉及。 - 取值范围：postPaid，后付费，即按需付费。 - 默认取值：不涉及。
	ChargeMode string `json:"charge_mode"`
}

func (o PostPaidChargeInfos) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostPaidChargeInfos struct{}"
	}

	return strings.Join([]string{"PostPaidChargeInfos", string(data)}, " ")
}
