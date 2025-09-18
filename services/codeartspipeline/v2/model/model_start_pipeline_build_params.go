package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StartPipelineBuildParams struct {

	// **参数解释**： 构建参数名。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Name string `json:"name"`

	// **参数解释**： 构建参数值。 **约束限制**： 不涉及。 **取值范围**： 不超过8192个字符。 **默认取值**： 不涉及。
	Value string `json:"value"`
}

func (o StartPipelineBuildParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartPipelineBuildParams struct{}"
	}

	return strings.Join([]string{"StartPipelineBuildParams", string(data)}, " ")
}
