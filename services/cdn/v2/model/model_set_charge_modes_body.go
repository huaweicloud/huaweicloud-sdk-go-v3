package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetChargeModesBody 设置计费模式请求体
type SetChargeModesBody struct {

	// **参数解释：** 计费模式 **约束限制：** v2及以上客户支持bw（带宽）计费模式 **取值范围：** - flux：流量 - bw：带宽 **默认取值：** 不涉及
	ChargeMode string `json:"charge_mode"`

	// **参数解释：** 产品模式 **约束限制：** 不涉及 **取值范围：** base：基础加速 **默认取值：** 不涉及
	ProductType string `json:"product_type"`

	// **参数解释：** 服务范围 **约束限制：** 不涉及 **取值范围：** mainland_china：中国大陆 **默认取值：** 不涉及
	ServiceArea string `json:"service_area"`
}

func (o SetChargeModesBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetChargeModesBody struct{}"
	}

	return strings.Join([]string{"SetChargeModesBody", string(data)}, " ")
}
