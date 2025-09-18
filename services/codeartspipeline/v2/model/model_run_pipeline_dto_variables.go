package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RunPipelineDtoVariables struct {

	// **参数解释**： 参数名。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Name string `json:"name"`

	// **参数解释**： 参数值。 **约束限制**： 不涉及。 **取值范围**： 不超过8192字符。 **默认取值**： 不涉及。
	Value string `json:"value"`
}

func (o RunPipelineDtoVariables) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunPipelineDtoVariables struct{}"
	}

	return strings.Join([]string{"RunPipelineDtoVariables", string(data)}, " ")
}
