package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateResourceTag struct {

	// **参数解释** 标签键。 **约束限制** 同一资源的key值不能重复。 最大长度36个字符。 **取值范围** 字符集：A-Z，a-z ， 0-9，‘-’，‘_’，UNICODE字符（\\u4E00-\\u9FFF）。 **默认取值** 不涉及。
	Key string `json:"key"`

	// **参数解释** 标签值。 **约束限制** 最大长度43个字符。 **取值范围** 字符集：A-Z，a-z ， 0-9，‘.’，‘-’，‘_’，UNICODE字符（\\u4E00-\\u9FFF）。 **默认取值** 不涉及。
	Value string `json:"value"`
}

func (o CreateResourceTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResourceTag struct{}"
	}

	return strings.Join([]string{"CreateResourceTag", string(data)}, " ")
}
