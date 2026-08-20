package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetChargeModesResponse Response Object
type SetChargeModesResponse struct {

	// **参数解释：** 账号的计费模式 **取值范围：** - flux：流量 - bw：带宽
	ChargeMode *string `json:"charge_mode,omitempty"`

	// **参数解释：** 加速类型 **取值范围：** base：基础加速
	ProductType *string `json:"product_type,omitempty"`

	// **参数解释：** 该模式生效时间 **取值范围：** 不涉及
	EffectiveTime *int64 `json:"effective_time,omitempty"`

	// **参数解释：** 创建时间 **取值范围：** 不涉及
	CreateTime *int64 `json:"create_time,omitempty"`

	// **参数解释：** 该模式的区域 **取值范围：** mainland_china：中国大陆
	ServiceArea *string `json:"service_area,omitempty"`

	// **参数解释：** 状态 > 首次开通状态为active（已生效）,之后修改为upcoming（待生效）  **取值范围：** - active：已生效 - upcoming：待生效
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SetChargeModesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetChargeModesResponse struct{}"
	}

	return strings.Join([]string{"SetChargeModesResponse", string(data)}, " ")
}
