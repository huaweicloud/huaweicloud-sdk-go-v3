package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RunPipelineDtoSources struct {

	// **参数解释**： 流水线源类型，目前支持“code”、“artifact”等代码源类型。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Type string `json:"type"`

	Params *RunPipelineDtoParams `json:"params,omitempty"`
}

func (o RunPipelineDtoSources) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunPipelineDtoSources struct{}"
	}

	return strings.Join([]string{"RunPipelineDtoSources", string(data)}, " ")
}
