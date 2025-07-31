package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IpsContent **参数解释**： 自定义IPS规则内容 **取值范围**： 不涉及
type IpsContent struct {

	// **参数解释**： 内容 **取值范围**： 不涉及
	Content *string `json:"content,omitempty"`

	// **参数解释**： 深度 **取值范围**： 不涉及
	Depth *int32 `json:"depth,omitempty"`

	// **参数解释**： 报文内容是否为十六进制 **取值范围**： 不涉及
	IsHex *bool `json:"is_hex,omitempty"`

	// **参数解释**： 是否忽略大小写 **取值范围**： 不涉及
	IsIgnore *bool `json:"is_ignore,omitempty"`

	// **参数解释**： 是否在uri中截取报文 **取值范围**： 不涉及
	IsUri *bool `json:"is_uri,omitempty"`

	// **参数解释**： 偏移量 **取值范围**： 不涉及
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**： 相对位置 **取值范围**： 不涉及
	RelativePosition *int32 `json:"relative_position,omitempty"`
}

func (o IpsContent) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IpsContent struct{}"
	}

	return strings.Join([]string{"IpsContent", string(data)}, " ")
}
