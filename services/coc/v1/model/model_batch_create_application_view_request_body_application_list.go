package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchCreateApplicationViewRequestBodyApplicationList struct {

	// **参数解释：** 应用名称。 **约束限制：** 不涉及。 **取值范围：** 由中文、英文字母、数字、中划线、下划线组成，长度在3~50个字符之间。 **默认取值：** 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 应用描述。 **约束限制：** 不涉及。 **取值范围：** 字符串，长度在0到256个字符之间。 **默认取值：** 不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释：** 父节点名称。 **约束限制：** 不涉及。 **取值范围：** 字符串，长度3到50个字符。 **默认取值：** 不涉及。
	ParentName *string `json:"parent_name,omitempty"`

	// **参数解释：** 层级，默认应用层级为1，子应用层级为2。 **约束限制：** 不涉及。 **取值范围：** - 1：应用层级。 - 2：子应用层级。 **默认取值：** 默认应用层级为1。
	Level *string `json:"level,omitempty"`
}

func (o BatchCreateApplicationViewRequestBodyApplicationList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateApplicationViewRequestBodyApplicationList struct{}"
	}

	return strings.Join([]string{"BatchCreateApplicationViewRequestBodyApplicationList", string(data)}, " ")
}
