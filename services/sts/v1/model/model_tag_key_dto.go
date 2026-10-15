package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TagKeyDto **参数解释**： 标签键。  **约束限制**： 长度为1到128个字符。只允许大小写字母、数字、空格、\"_\"、\".\"、\":\"、\"=\"、\"+\"、\"-\"、\"@\"、\"/\"，且不能以 \"&#95;sys&#95;\" 开头。  **取值范围**： 不涉及。  **默认取值**： 不涉及。
type TagKeyDto struct {
}

func (o TagKeyDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagKeyDto struct{}"
	}

	return strings.Join([]string{"TagKeyDto", string(data)}, " ")
}
