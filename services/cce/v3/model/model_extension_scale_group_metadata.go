package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExtensionScaleGroupMetadata 扩容伸缩组的基本信息
type ExtensionScaleGroupMetadata struct {

	// **参数解释**： 扩展伸缩组的uuid **约束限制**： - 创建节点池时自动生成，填写无效。 - 更新节点池时，如果填写则更新指定伸缩组。 - 更新节点池时，如果不填写则删除当前绑定的伸缩组，并创建新的伸缩组。  **取值范围**： 不涉及 **默认取值**： 不涉及
	Uid *string `json:"uid,omitempty"`

	// **参数解释**： 扩展伸缩组的名称。 **约束限制**： 不能为 **default**。 **取值范围**： 长度不能超过56个字符，只能包含数字和小写字母以及连字符（-），必须以小写字母开头以小写字母或者数字结尾。 **默认取值**： 不涉及
	Name *string `json:"name,omitempty"`
}

func (o ExtensionScaleGroupMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtensionScaleGroupMetadata struct{}"
	}

	return strings.Join([]string{"ExtensionScaleGroupMetadata", string(data)}, " ")
}
