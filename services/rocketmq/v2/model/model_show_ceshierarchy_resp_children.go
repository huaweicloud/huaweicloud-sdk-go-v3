package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCeshierarchyRespChildren **参数解释**： 子维度信息。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
type ShowCeshierarchyRespChildren struct {

	// **参数解释**： 子维度名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 监控指标名称列表。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Metrics *[]string `json:"metrics,omitempty"`

	// **参数解释**： 监控查询使用的key。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	KeyName *[]string `json:"key_name,omitempty"`

	// **参数解释**： 监控维度路由。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	DimRouter *[]string `json:"dim_router,omitempty"`
}

func (o ShowCeshierarchyRespChildren) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCeshierarchyRespChildren struct{}"
	}

	return strings.Join([]string{"ShowCeshierarchyRespChildren", string(data)}, " ")
}
