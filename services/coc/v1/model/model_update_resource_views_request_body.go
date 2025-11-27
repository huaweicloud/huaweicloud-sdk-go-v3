package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateResourceViewsRequestBody struct {

	// **参数解释：** 组织单元id。 **约束限制：** 不涉及。 **取值范围：** 自定义，视图所聚合的组织单元id列表。 **默认取值：** 不涉及。
	OrganizationUnitIds []string `json:"organization_unit_ids"`

	// **参数解释：** 资源类型列表。 **约束限制：** 不涉及。 **取值范围：** 自定义，用户创建视图时，选择的资源，资源对应类别组合成资源类型列表。 **默认取值：** 不涉及
	ResourceTypes []string `json:"resource_types"`

	// **参数解释：** 组织ID。 **约束限制：** 不涉及。 **取值范围：** 视图归属的组织id。 **默认取值：** 不涉及。
	OrganizationId *string `json:"organization_id,omitempty"`

	// **参数解释：** 视图类型。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	ViewType string `json:"view_type"`

	// **参数解释：** 视图名称。 **约束限制：** 不涉及。 **取值范围：** 用户自定义编辑。 **默认取值：** 不涉及。
	Name string `json:"name"`
}

func (o UpdateResourceViewsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateResourceViewsRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateResourceViewsRequestBody", string(data)}, " ")
}
