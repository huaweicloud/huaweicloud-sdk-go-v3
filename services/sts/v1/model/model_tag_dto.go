package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TagDto 标签，包含key、value两个字段。
type TagDto struct {

	// **参数解释**： 标签键。  **约束限制**： 不涉及。  **取值范围**： 长度为1到128个字符。只允许大小写字母、数字、空格、\"_\"、\".\"、\":\"、\"=\"、\"+\"、\"-\"、\"@\"、\"/\"，且不能以 \"&#95;sys&#95;\" 开头。  **默认取值**： 不涉及。
	Key string `json:"key"`

	// **参数解释**： 标签值。  **约束限制**： 不涉及。  **取值范围**： 长度为0到255个字符。只允许大小写字母、数字、空格、\"_\"、\".\"、\":\"、\"/\"、\"=\"、\"+\"、\"-\"、\"@\"。取值可以为空字符串，不可以为null。  **默认取值**： 不涉及。
	Value string `json:"value"`
}

func (o TagDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagDto struct{}"
	}

	return strings.Join([]string{"TagDto", string(data)}, " ")
}
