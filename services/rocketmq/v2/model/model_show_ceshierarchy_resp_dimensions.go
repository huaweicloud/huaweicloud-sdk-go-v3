package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowCeshierarchyRespDimensions struct {

	// **参数解释**： 监控维度名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 监控指标名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Metrics *[]string `json:"metrics,omitempty"`

	// **参数解释**： 监控查询使用的key。   **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	KeyName *[]string `json:"key_name,omitempty"`

	// **参数解释**： 监控维度路由。     **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	DimRouter *[]string `json:"dim_router,omitempty"`

	// **参数解释**： 子维度列表。     **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Children *[]ShowCeshierarchyRespChildren `json:"children,omitempty"`
}

func (o ShowCeshierarchyRespDimensions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCeshierarchyRespDimensions struct{}"
	}

	return strings.Join([]string{"ShowCeshierarchyRespDimensions", string(data)}, " ")
}
