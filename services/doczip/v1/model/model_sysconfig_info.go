package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SysconfigInfo 系统设置信息
type SysconfigInfo struct {

	// **参数解释**： 默认字体。 **约束限制**： 不涉及。 **取值范围**： [1-32]个字符。 **默认取值**： 不涉及。
	FontFamily *string `json:"font_family,omitempty"`

	// **参数解释**： 字体大小。 **约束限制**： 不涉及。 **取值范围**： [1-10]个字符。 **默认取值**： 不涉及。
	FontSize *string `json:"font_size,omitempty"`

	// **参数解释**： 行间距。 **约束限制**： 不涉及。 **取值范围**： [1-10]个字符。 **默认取值**： 不涉及。
	LineHeight *string `json:"line_height,omitempty"`
}

func (o SysconfigInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SysconfigInfo struct{}"
	}

	return strings.Join([]string{"SysconfigInfo", string(data)}, " ")
}
