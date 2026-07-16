package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PeriodOrderParam struct {

	// **参数解释**： 是否自动支付订单费用 **约束限制**： 不涉及 **取值范围**： - true：自动支付 - false：手动支付 **默认取值**： false
	IsAutoPay *bool `json:"isAutoPay,omitempty"`

	// **参数解释**： 是否自动续费 **约束限制**： 不涉及 **取值范围**： - true：自动续费 - false：不自动续费 **默认取值**： false
	IsAutoRenew *bool `json:"isAutoRenew,omitempty"`

	// **参数解释**： 包周期时间长度，不同局点、不同产品规格支持的范围可能不同，大部分局点支持：1-9月，1-3年，具体以接口返回为准。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	PeriodNum int32 `json:"periodNum"`

	// **参数解释**： 包周期单位 **约束限制**： 不涉及 **取值范围**： - \"month\"：月 - \"year\"：年 **默认取值**： 不涉及
	PeriodType string `json:"periodType"`
}

func (o PeriodOrderParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PeriodOrderParam struct{}"
	}

	return strings.Join([]string{"PeriodOrderParam", string(data)}, " ")
}
