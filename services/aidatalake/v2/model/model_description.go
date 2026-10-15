package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Description **参数解释**：描述信息。 **约束限制**：不涉及。 **取值范围**：长度为0-1024的任意字符组合，不包含“!”、“<”、“>”、“=”、“&”、“\"”、“'”。 **默认取值**：不涉及。
type Description struct {
}

func (o Description) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Description struct{}"
	}

	return strings.Join([]string{"Description", string(data)}, " ")
}
