package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExtensionScaleGroupMetadata 扩容伸缩组的基本信息
type ExtensionScaleGroupMetadata struct {

	// 扩展伸缩组的uuid，由系统自动生成
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
