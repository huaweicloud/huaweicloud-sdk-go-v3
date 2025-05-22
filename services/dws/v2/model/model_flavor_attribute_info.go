package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FlavorAttributeInfo **参数解释**： 扩展信息。 **取值范围**： 不涉及。
type FlavorAttributeInfo struct {

	// **参数解释**： 属性编码。 **取值范围**： 不涉及。
	Code string `json:"code"`

	// **参数解释**： 属性值。 **取值范围**： 不涉及。
	Value string `json:"value"`
}

func (o FlavorAttributeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlavorAttributeInfo struct{}"
	}

	return strings.Join([]string{"FlavorAttributeInfo", string(data)}, " ")
}
