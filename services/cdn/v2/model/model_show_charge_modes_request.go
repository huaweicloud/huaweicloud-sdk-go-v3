package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowChargeModesRequest Request Object
type ShowChargeModesRequest struct {

	// **参数解释：** 加速类型 **约束限制：** 不涉及 **取值范围：** - base：基础加速 **默认取值：** 不涉及
	ProductType string `json:"product_type"`

	// **参数解释：** 查询计费模式状态 **约束限制：** 不涉及 **取值范围：** - active：已生效 - upcoming：待生效 **默认取值：** active：已生效
	Status *string `json:"status,omitempty"`

	// **参数解释：** 服务范围 **约束限制：** 不涉及 **取值范围：** - mainland_china：中国大陆 - outside_mainland_china：中国大陆境外 **默认取值：** mainland_china：中国大陆
	ServiceArea *string `json:"service_area,omitempty"`
}

func (o ShowChargeModesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowChargeModesRequest struct{}"
	}

	return strings.Join([]string{"ShowChargeModesRequest", string(data)}, " ")
}
