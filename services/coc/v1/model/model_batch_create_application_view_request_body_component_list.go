package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchCreateApplicationViewRequestBodyComponentList struct {

	// **参数解释：** 组件名称。 **约束限制：** 不涉及。 **取值范围：** 由中文、英文字母、数字、中划线、下划线组成，长度在3~50个字符之间。 **默认取值：** 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 父节点名称。 **约束限制：** 不涉及。 **取值范围：** 字符串，长度3到50个字符。 **默认取值：** 不涉及。
	ParentName *string `json:"parent_name,omitempty"`
}

func (o BatchCreateApplicationViewRequestBodyComponentList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateApplicationViewRequestBodyComponentList struct{}"
	}

	return strings.Join([]string{"BatchCreateApplicationViewRequestBodyComponentList", string(data)}, " ")
}
