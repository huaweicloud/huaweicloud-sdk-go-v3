package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IpsContent **参数解释**： 自定义IPS规则内容 **约束限制**：     不涉及 **取值范围**： 不涉及 **默认取值**：   不涉及
type IpsContent struct {

	// **参数解释**： 内容 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**：   不涉及
	Content *string `json:"content,omitempty"`

	// **参数解释**： 匹配特征时，截止匹配的位置 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**：   不涉及
	Depth *int32 `json:"depth,omitempty"`

	// **参数解释**： 报文内容是否为十六进制 **约束限制**： 不涉及 **取值范围**： true： 是十六进制 false： 不是十六进制 **默认取值**：   不涉及
	IsHex *bool `json:"is_hex,omitempty"`

	// **参数解释**： 是否忽略大小写 **约束限制**： 不涉及 **取值范围**： true： 忽略 false： 不忽略 **默认取值**：   不涉及
	IsIgnore *bool `json:"is_ignore,omitempty"`

	// **参数解释**： 是否匹配URL中跟“内容”一致的字段 **约束限制**： 不涉及 **取值范围**： true： 匹配 false： 不匹配 **默认取值**：   不涉及
	IsUri *bool `json:"is_uri,omitempty"`

	// **参数解释**： 匹配特征时开始的位置 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**：   不涉及
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**： 匹配特征时，指定开始的位置 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**：   不涉及
	RelativePosition *int32 `json:"relative_position,omitempty"`
}

func (o IpsContent) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IpsContent struct{}"
	}

	return strings.Join([]string{"IpsContent", string(data)}, " ")
}
