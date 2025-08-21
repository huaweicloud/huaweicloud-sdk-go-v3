package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchCreateApplicationViewRequestBody struct {

	// **参数解释：** 应用信息组成的列表。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	ApplicationList []BatchCreateApplicationViewRequestBodyApplicationList `json:"application_list"`

	// **参数解释：** 组件信息组成的列表。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	ComponentList *[]BatchCreateApplicationViewRequestBodyComponentList `json:"component_list,omitempty"`

	// **参数解释：** 分组信息组成的列表。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	GroupList *[]BatchCreateApplicationViewRequestBodyGroupList `json:"group_list,omitempty"`
}

func (o BatchCreateApplicationViewRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateApplicationViewRequestBody struct{}"
	}

	return strings.Join([]string{"BatchCreateApplicationViewRequestBody", string(data)}, " ")
}
