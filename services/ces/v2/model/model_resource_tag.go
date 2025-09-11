package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceTag **参数解释**： 资源标签。   **约束限制**： 不涉及。
type ResourceTag struct {

	// **参数解释**： 标签名。 **约束限制**： 不涉及。 **取值范围**： 最大长度128个unicode字符。           **默认取值**： 不涉及。
	Key string `json:"key"`

	// **参数解释**： 标签值。 **约束限制**： 不涉及。 **取值范围**： 最大长度255个unicode字符。           **默认取值**： 不涉及。
	Value string `json:"value"`
}

func (o ResourceTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceTag struct{}"
	}

	return strings.Join([]string{"ResourceTag", string(data)}, " ")
}
